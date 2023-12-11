package core

// Implementation of (*State).Dump()

import (
	"encoding/json"
	"time"

	"github.com/neondatabase/autoscaling/pkg/api"
)

func shallowCopy[T any](ptr *T, f ...func(T) T) *T {
	if ptr == nil {
		return nil
	} else {
		x := *ptr
		if len(f) != 0 {
			x = f[0](x)
		}
		return &x
	}
}

// StateDump provides introspection into the current values of the fields of State
//
// It implements json.Marshaler.
type StateDump struct {
	internal internalState
}

func (d StateDump) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.internal)
}

// Dump produces a JSON-serializable copy of the State
func (s *State) Dump() StateDump {
	return StateDump{
		internal: internalState{
			Debug:   s.internal.Debug,
			Config:  s.internal.Config,
			VM:      s.internal.VM,
			Plugin:  s.internal.Plugin.deepCopy(),
			Monitor: s.internal.Monitor.deepCopy(),
			NeonVM:  s.internal.NeonVM.deepCopy(),
			Metrics: shallowCopy[api.Metrics](s.internal.Metrics),
		},
	}
}

func (s *pluginState) deepCopy() pluginState {
	return pluginState{
		OngoingRequest: s.OngoingRequest,
		ComputeUnit:    shallowCopy[api.Resources](s.ComputeUnit),
		LastRequest: shallowCopy(s.LastRequest, func(r pluginRequested) pluginRequested {
			return pluginRequested{At: r.At, Resources: r.Resources}
		}),
		LastFailureAt: shallowCopy[time.Time](s.LastFailureAt),
		Permit:        shallowCopy[api.Resources](s.Permit),
	}
}

func (s *monitorState) deepCopy() monitorState {
	return monitorState{
		OngoingRequest: shallowCopy(s.OngoingRequest, func(r ongoingMonitorRequest) ongoingMonitorRequest {
			return ongoingMonitorRequest{Kind: r.Kind, Requested: r.Requested}
		}),
		RequestedUpscale: shallowCopy(s.RequestedUpscale, func(r requestedUpscale) requestedUpscale {
			return requestedUpscale{At: r.At, Base: r.Base, Requested: r.Requested}
		}),
		DeniedDownscale: shallowCopy(s.DeniedDownscale, func(r deniedDownscale) deniedDownscale {
			return deniedDownscale{At: r.At, Current: r.Current, Requested: r.Requested}
		}),
		Approved:           shallowCopy[api.Resources](s.Approved),
		DownscaleFailureAt: shallowCopy[time.Time](s.DownscaleFailureAt),
		UpscaleFailureAt:   shallowCopy[time.Time](s.UpscaleFailureAt),
	}
}

func (s *neonvmState) deepCopy() neonvmState {
	return neonvmState{
		LastSuccess:      shallowCopy[api.Resources](s.LastSuccess),
		OngoingRequested: shallowCopy[api.Resources](s.OngoingRequested),
		RequestFailedAt:  shallowCopy[time.Time](s.RequestFailedAt),
	}
}
