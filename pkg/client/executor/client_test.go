package executor

import (
	"Color-FaaS-Core/pkg/common"
	pb "Color-FaaS-Core/pkg/proto/executor"
	"encoding/json"
	"testing"
)

func TestExecutorClient_Heartbeat(t *testing.T) {
	exe, _ := New("127.0.0.1", "50001")
	reply, _ := exe.Heartbeat()

	reqByte, _ := json.Marshal(reply)
	println(string(reqByte))
}

func TestExecutorClient_InitTask(t *testing.T) {
	exe, _ := New("127.0.0.1", "50001")

	req := pb.TaskInstance{
		TaskID:          "0",
		FuncName:        "helloworld",
		FuncID:          "0",
		FuncStorageType: common.HDFS,
		TaskFuncPath:    "/color-faas/helloworld_raw",
	}

	replyInit, _ := exe.InitTask(&req)
	reqByte, _ := json.Marshal(replyInit)
	println(string(reqByte))

	replyRun, _ := exe.RunTask(&req)
	reqByte, _ = json.Marshal(replyRun)
	println(string(reqByte))
}
