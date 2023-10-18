package main

import (
	"Color-FaaS-Core/pkg/executor"
	"log"
)

func main() {
	exe, err := executor.New()

	if err != nil {
		log.Fatalf("ERROR %v", err)
	}

	log.Default().Printf("Start Executor RPC server...")
	exe.Start()
}
