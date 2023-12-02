package configs

import (
	"Color-FaaS-Core/pkg/model"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const clientCfgPathDebug = "../../../conf/client_conf.yaml"
const clientCfgPathDev = "../conf/client_conf.yaml"

type CConfig struct {
	Port            string `yaml:"Port"`
	ExecutorPath    string `yaml:"ExecutorPath"`
	MaxCacheFuncNum int    `yaml:"MaxCacheFuncNum"`
	ZKPath          string `yaml:"ZKPath"`
}

type ClientConfig struct {
	Cfg CConfig `yaml:"ExecutorConfig"`
}

var defaultClientConfig = ClientConfig{
	Cfg: CConfig{
		Port:            "9090",
		ExecutorPath:    "../executor/executor",
		MaxCacheFuncNum: 50,
		ZKPath:          "localhost:2181",
	},
}

func NewClientConfig(runtimeInfo model.RuntimeInfo) ClientConfig {
	path := clientCfgPathDebug
	if !runtimeInfo.IsDebug {
		path = clientCfgPathDev
	}

	cfgByte, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error: %v", err)
		return defaultClientConfig
	}

	var cfg ClientConfig
	err = yaml.Unmarshal(cfgByte, &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
		return defaultClientConfig
	}
	log.Printf("config for executor :%v", cfg)
	return cfg
}
