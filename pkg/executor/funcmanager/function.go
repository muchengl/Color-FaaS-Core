package funcmanager

import common "Color-FaaS-Core/pkg/common"

type FunctionInstance struct {
	FuncName    string
	FuncID      string
	StorageType common.StorageType
	RemotePath  string

	LocalPath string
	Status    common.FuncRunningStatus
}
