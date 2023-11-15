package manager

import "Color-FaaS-Core/pkg/executor/env"

type pool interface {
	// download the func or renew func
	getFunc(instance *env.FunctionInstance) (*env.FunctionInstance, error)
}
