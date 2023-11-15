package env

import (
	"Color-FaaS-Core/pkg/common"
	"Color-FaaS-Core/pkg/model"
	"Color-FaaS-Core/pkg/utils"
	"errors"
	"log"
)

type Getter interface {
	DownloadFile(instance FunctionInstance) error
}

type FuncGetter struct {
	//hdfsCfg     hdfsCfg
	hdfsManager utils.HdfsManager
	//s3Cfg       s3Cfg
	s3Manager utils.S3Manager
}

func NewGetter(info model.RuntimeInfo) (*FuncGetter, error) {
	get := FuncGetter{}
	hdfsManager, err := utils.NewHdfsManager(info)
	if err != nil {
		return nil, err
	}
	get.hdfsManager = *hdfsManager

	// todo S3

	return &get, nil
}

func (g *FuncGetter) DownloadFile(instance FunctionInstance) error {
	if instance.StorageType == common.HDFS {
		err := g.downloadHDFSFile(instance)
		if err != nil {
			return err
		}
	}
	if instance.StorageType == common.S3 {
		err := g.downloadS3File(instance)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *FuncGetter) downloadHDFSFile(instance FunctionInstance) error {
	err := g.hdfsManager.DownloadFile(instance.RemotePath, instance.LocalPath)
	if err != nil {
		log.Printf("hdfs download fail: %v", err)
		return err
	}
	return nil
}

func (g *FuncGetter) downloadS3File(instance FunctionInstance) error {
	return errors.New("s3 err")
}
