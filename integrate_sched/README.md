# Integrated scheduler

This builds on [`scaling_builtin`], which in turn builds on [`kind_setup`]. For more detail, refer
to their respective READMEs (particularly [`kind_setup`]).

[`scaling_builtin`]: ../scaling_builtin
[`kind_setup`]: ../kind_setup

Steps:

Build everything:

```sh
vm_image/start-local-registry.sh # required for everything below. Does nothing on repeat
vm_image/build.sh
scheduler/build.sh
autoscaler-agent/build.sh
```

Download kubernetes dependencies:

```sh
./download-cni.sh
curl -sS https://raw.githubusercontent.com/flannel-io/flannel/v0.19.2/Documentation/kube-flannel.yml \
    -o flannel.yaml
curl -sSL https://github.com/cert-manager/cert-manager/releases/download/v1.8.2/cert-manager.yaml \
    -o cert-manager.yaml
curl -sS https://raw.githubusercontent.com/k8snetworkplumbingwg/multus-cni/master/deployments/multus-daemonset.yml \
    -o multus-daemonset.yaml
```

Set up the cluster:

```sh
kind create cluster -n autoscale-sched --config=kind-config.yaml
kubectl apply -f flannel.yaml -f cert-manager.yaml -f multus-daemonset.yaml \
    -f scheduler-deploy.yaml -f autoscaler-agent-deploy.yaml
# (wait until cert manager has been ready for a bit)
kubectl apply -f virtink_localhost:5001.yaml
```

Run the VM(s):

```sh
kubectl apply -f vm-deploy.yaml
# or:
kubectl apply -f vm-double-deploy.yaml
```

Run pgbench and watch the vCPU allocation grow:

```sh
./run-bench.sh
# or:
VM_NAME=postgres14-disk-1 ./run-bench.sh
VM_NAME=postgres14-disk-2 ./run-bench.sh
```

## Architecture

Broadly speaking, the components are:

* `autoscaler-agent`:
  * Gets metrics from VM's `node_exporter` (`http://10.0.2.2:9100/metrics`)
  * Uses `ch-remote` to communicate with `cloud-hypervisor` to resize vCPU (via `/var/run/virtink/ch.sock`)
  * Requests vCPU increase, notify vCPU decrease to the scheduler plugin (configured port `10299`)
* `scheduler`:
  * Provides the `AutoscaleEnforcer` scheduler plugin, using the `kube-scheduler` plugin interface
  * Handles requests / notifications from `autoscaler-agent`s, limiting them to the node's capacity.
      (Currently reserves 20% of node CPU capacity for system tasks.)
* `vm_image`:
  * The normal VM + `node_exporter` image that we've used in other places.
  * **NOTE:** The entire setup requires the registry at `localhost:5001` from `start-local-registry.sh`
* `virtink_localhost:5001`:
  * Autogenerated virtink deployment from https://github.com/neondatabase/autoscaling/tree/cpu-scaling
    at commit 9e1f8b4b.

Some basics on the `autoscaler-agent` \<-\> `scheduler` protocol:

* All messages (currently) are `autoscaler-agent` -> `scheduler`, sending a `ResourceRequest` and
  expecting a `Permit` as the response.
* When the `autoscaler-agent` wants to *decrease* vCPU, it does that immediately and the
    `ResourceRequest` is informative (i.e., "Hey scheduler, just so you know, I've decreased my
    vCPU. You can allocate that elsewhere").
* When the `autoscaler-agent` wants to *increase* vCPU, it submits the `ResourceRequest` and the
    scheduler returns a `Permit` that may be lower than the requested vCPU amount, but not lower
    than the current vCPU (i.e., "Hey scheduler, can I increase to `X` vCPUs?" ... "You can only go
    to `Y`", where `current <= Y <= X`).
  * When the scheduler responds to an increase request by not allowing *any* increase, we log that
      request as "denied" (in the `autoscaler-agent`).

The files implementing the protocol are in `autoscaler-agent/src/run.go` and `scheduler/src/run.go`.

Currently, the scheduler also appropriately handles pod un-scheduling via `Reserve`/`Unreserve`,
with some initial (but non-binding) capacity checks in the `Filter` step.

Also: the scheduler currently does not handle VM/pod deletion, so it will "leak" the space reserved
for deleted pods.