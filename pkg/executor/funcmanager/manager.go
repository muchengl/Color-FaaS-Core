package funcmanager

import (
	"Color-FaaS-Core/pkg/common"
	"Color-FaaS-Core/pkg/model"
	"errors"
	"log"
)

type FuncManager struct {
	funcPool pool

	info   model.RuntimeInfo
	getter funcGetter
}

func New(info model.RuntimeInfo) (*FuncManager, error) {
	mgr := FuncManager{
		info: info,
	}

	getter, err := newGetter(info)
	if err != nil {
		log.Printf("init function manager fail: %v", err)
		return nil, err
	}
	mgr.getter = *getter

	// use the getter in manager, so getter can be updated
	p, err := NewLruPool(getter)
	if err != nil {
		log.Printf("init function manager fail: %v", err)
		return nil, err
	}
	mgr.funcPool = p

	return &mgr, nil
}

func (f *FuncManager) Init(instance FunctionInstance) error {
	runnableFuncInstance, err := f.funcPool.getFunc(&instance)
	if err != nil {
		log.Printf("can't get FunctionInstance, %v", instance)
		return err
	}

	if runnableFuncInstance.Status != common.Init {
		log.Printf("can't get FunctionInstance, %v", instance)
		return errors.New("init FunctionInstance fail")
	}

	return nil
}

func (f *FuncManager) Start(instance FunctionInstance) error {
	runnableFuncInstance, err := f.funcPool.getFunc(&instance)
	if err != nil {
		log.Printf("can't get FunctionInstance, %v", instance)
		return err
	}

	err = runnableFuncInstance.run()
	if err != nil {
		log.Printf("run FunctionInstance fail, %v", instance)
		return err
	}
	return nil
}

func (f *FuncManager) Kill(instance FunctionInstance) error {
	runnableFuncInstance, err := f.funcPool.getFunc(&instance)
	if err != nil {
		return err
	}

	runnableFuncInstance.Status = common.Die
	// todo kill the func by Http
	return nil
}
