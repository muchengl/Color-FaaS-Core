package main

import (
	"Color-FaaS-Core/pkg/executor"
	"Color-FaaS-Core/pkg/model"
	"log"
	"os"
)

func main() {
	exe, err := executor.New()
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

	if err != nil {
		log.Fatalf("ERROR %v", err)
	}

	log.Default().Printf("Start Executor RPC server...")
	exe.Start()
}
