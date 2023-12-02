package main

import (
	client "Color-FaaS-Core/pkg/client"
	"Color-FaaS-Core/pkg/model"
	"log"
	"os"
)

func main() {
	exe, err := client.NewClient(model.DefaultInfo)
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

	log.Default().Printf("Start Client Http server...")
	exe.Run()
}
