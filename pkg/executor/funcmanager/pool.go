package funcmanager

type pool interface {
	// download the func or renew func
	getFunc(instance *FunctionInstance) (*FunctionInstance, error)
}
