package agent

import (
	"context"
	"fmt"

	"github.com/tychoish/fun/pubsub"
	"go.uber.org/zap"

	"k8s.io/client-go/kubernetes"

	vmclient "github.com/neondatabase/autoscaling/neonvm/client/clientset/versioned"
	"github.com/neondatabase/autoscaling/pkg/agent/billing"
	"github.com/neondatabase/autoscaling/pkg/agent/schedwatch"
	"github.com/neondatabase/autoscaling/pkg/util"
	"github.com/neondatabase/autoscaling/pkg/util/watch"
)

type MainRunner struct {
	EnvArgs    EnvArgs
	Config     *Config
	KubeClient *kubernetes.Clientset
	VMClient   *vmclient.Clientset
}

func (r MainRunner) Run(logger *zap.Logger, ctx context.Context) error {
	vmEventQueue := pubsub.NewUnlimitedQueue[vmEvent]()
	defer vmEventQueue.Close()
	pushToQueue := func(ev vmEvent) {
		if err := vmEventQueue.Add(ev); err != nil {
			logger.Warn("Failed to add vmEvent to queue", zap.Object("event", ev), zap.Error(err))
		}
	}

	watchMetrics := watch.NewMetrics("autoscaling_agent_watchers")

	perVMMetrics, vmPromReg := makePerVMMetrics()

	logger.Info("Starting VM watcher")
	vmWatchStore, err := startVMWatcher(ctx, logger, r.Config, r.VMClient, watchMetrics, perVMMetrics, r.EnvArgs.K8sNodeName, pushToQueue)
	if err != nil {
		return fmt.Errorf("Error starting VM watcher: %w", err)
	}
	defer vmWatchStore.Stop()
	logger.Info("VM watcher started")

	schedTracker, err := schedwatch.StartSchedulerWatcher(ctx, logger, r.KubeClient, watchMetrics, r.Config.Scheduler.SchedulerName)
	if err != nil {
		return fmt.Errorf("Starting scheduler watch server: %w", err)
	}
	defer schedTracker.Stop()

	globalState, globalPromReg := r.newAgentState(logger, r.EnvArgs.K8sPodIP, schedTracker)
	watchMetrics.MustRegister(globalPromReg)

	logger.Info("Starting billing metrics collector")
	storeForNode := watch.NewIndexedStore(vmWatchStore, billing.NewVMNodeIndex(r.EnvArgs.K8sNodeName))

	metrics := billing.NewPromMetrics()
	metrics.MustRegister(globalPromReg)

	// TODO: catch panics here, bubble those into a clean-ish shutdown.
	go billing.RunBillingMetricsCollector(ctx, logger, &r.Config.Billing, storeForNode, metrics)

	promLogger := logger.Named("prometheus")
	if err := util.StartPrometheusMetricsServer(ctx, promLogger.Named("global"), 9100, globalPromReg); err != nil {
		return fmt.Errorf("Error starting prometheus metrics server: %w", err)
	}
	if err := util.StartPrometheusMetricsServer(ctx, promLogger.Named("per-vm"), 9101, vmPromReg); err != nil {
		return fmt.Errorf("Error starting prometheus metrics server: %w", err)
	}

	if r.Config.DumpState != nil {
		logger.Info("Starting 'dump state' server")
		if err := globalState.StartDumpStateServer(ctx, logger.Named("dump-state"), r.Config.DumpState); err != nil {
			return fmt.Errorf("Error starting dump state server: %w", err)
		}
	}

	logger.Info("Entering main loop")
	for {
		event, err := vmEventQueue.Wait(ctx)
		if err != nil {
			if ctx.Err() != nil {
				// treat context canceled as a "normal" exit (because it is)
				return nil
			}

			logger.Error("vmEventQueue returned error", zap.Error(err))
			return err
		}
		globalState.handleEvent(ctx, logger, event)
	}
}
