package clientmodel

type TaskInstance struct {
	TaskID          string `json:"taskID"`
	FuncName        string `json:"funcName"`
	FuncID          string `json:"funcID"`
	FuncStorageType string `json:"funcStorageType"` // s3
	TaskFuncPath    string `json:"taskFuncPath"`
	TaskRunningMode string `json:"taskRunningMode"` //sync or async
	FuncType        string `json:"funcType"`        // Go or python ......
	TaskCPUCore     int64  `json:"taskCPUCore"`
	TaskMem         int64  `json:"taskMem"`
	TaskDiskSpace   int64  `json:"taskDiskSpace"`
	TaskMaxTime     int64  `json:"taskMaxTime"`
	TaskInput       string `json:"taskInput"`
}
