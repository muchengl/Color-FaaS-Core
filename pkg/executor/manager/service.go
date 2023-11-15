package manager

import "Color-FaaS-Core/pkg/executor/env"

type Service interface {
	Init(instance env.FunctionInstance) error
	Start(instance env.FunctionInstance) error
	Kill(instance env.FunctionInstance) error
}
