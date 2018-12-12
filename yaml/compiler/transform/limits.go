package transform

import "github.com/drone/drone-runtime/engine"
// WithLimits is a transform function that applies
// resource limits to the container processes.
func WithLimits(memlimit int64, cpulimit int64) func(*engine.Spec) {
	return func(spec *engine.Spec) {
		// if no limits are defined exit immediately.
		if memlimit == 0 && cpulimit == 0 {
			return
		}
		// otherwise apply the resource limits to every
		// step in the runtime spec.
		for _, step := range spec.Steps {
			if step.Resources == nil {
				step.Resources = &engine.Resources{}
			}
			if step.Resources.Limits == nil {
				step.Resources.Limits = &engine.ResourceObject{}
			}
			if memlimit != 0 {
				step.Resources.Limits.Memory = memlimit
			}
			if cpulimit != 0 {
				step.Resources.Limits.CPU = cpulimit * 1000
			}
		}
	}
}
