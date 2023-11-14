package funcmanager

import (
	common "Color-FaaS-Core/pkg/common"
	configs "Color-FaaS-Core/pkg/configs"
	pb "Color-FaaS-Core/pkg/proto/executor"
	"errors"
	"github.com/google/uuid"
	"os/exec"
)

type FunctionInstance struct {
	FuncName    string
	FuncID      string
	StorageType common.StorageType
	RemotePath  string

	LocalPath string
	Status    common.FuncRunningStatus
}

func (f *FunctionInstance) Init(req *pb.TaskInstance, cfg configs.ExecutorConfig) {
	f.FuncName = req.FuncName
	f.FuncID = req.FuncID
	f.StorageType = req.FuncStorageType
	f.RemotePath = req.TaskFuncPath
	f.LocalPath = cfg.Cfg.FuncFilePath + "/" + uuid.NewString()
	f.Status = common.NotInit
}

func (f *FunctionInstance) run() error {
	if f.Status != common.Init {
		return errors.New("FunctionInstance not init, can't run")
	}

	// todo : should use runtime, like cgroups, can't run it directly
	cmd := exec.Command(f.LocalPath)
	err := cmd.Start()
	if err != nil {
		return err
	}

	f.Status = common.Alive
	return nil
}
