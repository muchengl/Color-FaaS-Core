package executor

import (
	"encoding/json"
	"testing"
)

func TestExecutorClient_Heartbeat(t *testing.T) {
	exe, _ := New("127.0.0.1", "50001")
	reply, _ := exe.Heartbeat()

	reqByte, _ := json.Marshal(reply)
	println(string(reqByte))
}
