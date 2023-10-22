package utils

import (
	"Color-FaaS-Core/pkg/model"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Test_New(t *testing.T) {
	mgr, err := NewHdfsManager(model.DefaultInfo)

	assert.Equal(t, nil, err)
	log.Print(mgr)
}

func Test_Download(t *testing.T) {
	// waiting for e2e test

	//mgr, _ := NewHdfsManager(model.DefaultInfo)
	//mgr.DownloadFile("/color-faas/helloworld_raw", "./")
}
