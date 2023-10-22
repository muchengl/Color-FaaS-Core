package executor

import (
	"Color-FaaS-Core/pkg/model"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Test_New(t *testing.T) {
	cfg := newConfig(model.DefaultInfo)

	assert.NotEqual(t, nil, cfg)
	assert.NotEqual(t, nil, cfg.Cfg.FuncRunDir)
	assert.NotEqual(t, nil, cfg.Cfg.FuncFilePath)
	log.Print(cfg)
}
