package executor

import (
	"log"
	"os"

	"Color-FaaS-Core/pkg/model"
	"gopkg.in/yaml.v2"
)

const cfgPathDebug = "../../conf/executor_conf.yaml"
const cfgPathDev = "../conf/executor_conf.yaml"

type ExecutorConfig struct {
	Port         string `yaml:"Port"`
	FuncFilePath string `yaml:"FuncFilePath"`
	FuncRunDir   string `yaml:"FuncRunDir"`
}

type config struct {
	Cfg ExecutorConfig `yaml:"ExecutorConfig"`
}

var defaultConfig = config{
	Cfg: ExecutorConfig{
		Port:         "50001",
		FuncFilePath: "./funcs",
		FuncRunDir:   "./run",
	},
}

func newConfig(runtimeInfo model.RuntimeInfo) config {
	path := cfgPathDebug
	if !runtimeInfo.IsDebug {
		path = cfgPathDev
	}

	cfgByte, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error: %v", err)
		return defaultConfig
	}

	var cfg config
	err = yaml.Unmarshal(cfgByte, &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
		return defaultConfig
	}
	log.Printf("config for executor :%v", cfg)
	return cfg
}
