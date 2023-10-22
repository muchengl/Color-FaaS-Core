package common

type StorageType = string

const (
	HDFS StorageType = "hdfs"
	S3   StorageType = "awsS3"
)

type FuncRunningMode = string

const (
	Sync  FuncRunningMode = "sync"
	Async FuncRunningMode = "async"
)

type FuncRunningStatus = string

const (
	// NotInit func's file has not been downloaded to local
	NotInit FuncRunningStatus = "notinit"

	// Init func's file has been downloaded to local
	Init FuncRunningStatus = "init"

	// Alive func is alive, waiting for task submit
	Alive FuncRunningStatus = "running"

	// Running func is running a task
	Running FuncRunningStatus = "pending"

	// Offline func's heartbeat msg lost
	Offline FuncRunningStatus = "offline"

	// Die func already been killed
	Die FuncRunningStatus = "die"
)
