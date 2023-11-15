package configs

import (
	"log"
	"os"

	"Color-FaaS-Core/pkg/model"
	"gopkg.in/yaml.v2"
)

const cfgPathDebug = "../../../conf/executor_conf.yaml"
const cfgPathDev = "../conf/executor_conf.yaml"

type EConfig struct {
	Port            string `yaml:"Port"`
	FuncFilePath    string `yaml:"FuncFilePath"`
	FuncRunDir      string `yaml:"FuncRunDir"`
	MaxCacheFuncNum int    `yaml:"MaxCacheFuncNum"`
}

type ExecutorConfig struct {
	Cfg EConfig `yaml:"ExecutorConfig"`
}

var defaultExecutorConfig = ExecutorConfig{
	Cfg: EConfig{
		Port:            "50001",
		FuncFilePath:    "./funcs",
		FuncRunDir:      "./run",
		MaxCacheFuncNum: 50,
	},
}

func NewConfig(runtimeInfo model.RuntimeInfo) ExecutorConfig {
	path := cfgPathDebug
	if !runtimeInfo.IsDebug {
		path = cfgPathDev
	}

	cfgByte, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error: %v", err)
		return defaultExecutorConfig
	}

	var cfg ExecutorConfig
	err = yaml.Unmarshal(cfgByte, &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
		return defaultExecutorConfig
	}
	log.Printf("config for executor :%v", cfg)
	return cfg
}
