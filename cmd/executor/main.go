package main

import (
	"Color-FaaS-Core/pkg/executor"
	"Color-FaaS-Core/pkg/model"
	"log"
	"os"
)

func main() {
	exe, err := executor.New(model.DefaultInfo)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	exe.RuntimeInfo = model.DefaultInfo
	err = exe.RuntimeInfo.InitByArgs(os.Args)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	err = exe.InitRunningEnv()
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	log.Default().Printf("Start Executor RPC server...")
	exe.Start()
}
