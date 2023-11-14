package funcmanager

type Service interface {
	Init(instance FunctionInstance) error
	Start(instance FunctionInstance) error
	Kill(instance FunctionInstance) error
}
