package funcmanager

import (
	"Color-FaaS-Core/pkg/common"
	"Color-FaaS-Core/pkg/model"
	"log"
	"os/exec"
)

type FuncManager struct {
	// functionLruPool
	function FunctionInstance
	info     model.RuntimeInfo
	getter   funcGetter
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

	return &mgr, nil
}

func (f *FuncManager) Init(instance FunctionInstance) error {
	f.function = instance

	// download func file
	err := f.getter.downloadFile(instance)
	if err != nil {
		log.Printf("download function fail: %v", err)
		return err
	}
	instance.Status = common.Init
	return nil
}

func (f *FuncManager) Start() error {

	// todo : should use runtime, can't run it directly
	cmd := exec.Command(f.function.LocalPath)
	cmd.Start()

	f.function.Status = common.Alive
	return nil
}

func (f *FuncManager) Kill() error {

	f.function.Status = common.Die
	return nil
}
