package funcmanager

type Service interface {
	Init(instance FunctionInstance) error
	Start() error
	Kill() error
}
