package utils

import (
	"Color-FaaS-Core/pkg/model"
	"errors"
	"fmt"
	"github.com/colinmarc/hdfs/v2"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
)

const cfgPathDebug = "../../conf/hdfs_conf.yaml"
const cfgPathDev = "../conf/hdfs_conf.yaml"

type HdfsConfig struct {
	Config struct {
		NameNodeHost string `yaml:"NameNodeHost"`
		NameNodePort string `yaml:"NameNodePort"`
	} `yaml:"HdfsConfig"`
}

type HdfsManager struct {
	cfg         HdfsConfig
	RuntimeInfo model.RuntimeInfo
	client      *hdfs.Client
}

func NewHdfsManager(info model.RuntimeInfo) (*HdfsManager, error) {
	mgr := HdfsManager{}
	mgr.RuntimeInfo = info

	// init config
	if info.CfgType == model.Local {
		cfgPath := cfgPathDebug
		if !info.IsDebug {
			cfgPath = cfgPathDev
		}
		cfgByte, err := os.ReadFile(cfgPath)
		if err != nil {
			log.Fatalf("error: %v", err)
			return nil, err
		}

		var config HdfsConfig
		err = yaml.Unmarshal(cfgByte, &config)
		if err != nil {
			log.Fatalf("error: %v", err)
			return nil, err
		}
		log.Printf("config for hdfs :%v", config)
		mgr.cfg = config
	} else {
		log.Printf("no support config: %v", info.CfgType)
		return nil, errors.New("not supported config" + string(info.CfgType))
	}

	// HDFS Client
	url := mgr.cfg.Config.NameNodeHost + ":" + mgr.cfg.Config.NameNodePort
	hdfsClient, err := hdfs.New(url)
	if err != nil {
		fmt.Printf("Failed to create client: %s\n", err)
		return nil, err
	}
	mgr.client = hdfsClient

	return &mgr, nil
}

func (h *HdfsManager) DownloadFile(hdfsPath string, localPath string) error {
	file, err := h.client.Open(hdfsPath)
	if err != nil {
		log.Printf("Failed to open file from HDFS: %s\n", err)
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Failed to read from HDFS file: %s\n", err)
		return err
	}

	err = os.WriteFile(localPath, data, 0777)
	if err != nil {
		log.Printf("Failed to write to local file: %s\n", err)
		return err
	}

	log.Printf("File downloaded successfully:%s -> %s", hdfsPath, localPath)
	return nil
}
