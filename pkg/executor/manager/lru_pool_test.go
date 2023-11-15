package manager

import (
	"Color-FaaS-Core/pkg/common"
	"Color-FaaS-Core/pkg/configs"
	"Color-FaaS-Core/pkg/executor/env"
	mock "Color-FaaS-Core/pkg/mock/executor/funcmanager"
	"Color-FaaS-Core/pkg/model"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getFunc_Len(t *testing.T) {
	info := model.RuntimeInfo{
		CfgType: model.Local,
		IsDebug: true,
	}

	cfg := configs.NewConfig(info)

	ctrl := gomock.NewController(t)
	getter := mock.NewMockGetter(ctrl)
	getter.EXPECT().DownloadFile(gomock.Any()).AnyTimes().Return(nil)

	pool, _ := NewLruPool(getter, &cfg)

	pool.cfg.Cfg.MaxCacheFuncNum = 10

	for i := 0; i < 15; i++ {
		println("get:", i)
		instance := env.FunctionInstance{
			FuncID:      fmt.Sprint(i),
			StorageType: common.HDFS,
			RemotePath:  "test",
		}
		pool.getFunc(&instance)
	}

	assert.Equal(t, 10, pool.cacheQueue.Len())
}

func Test_getFunc_Element(t *testing.T) {
	info := model.RuntimeInfo{
		CfgType: model.Local,
		IsDebug: true,
	}

	cfg := configs.NewConfig(info)

	ctrl := gomock.NewController(t)
	getter := mock.NewMockGetter(ctrl)
	getter.EXPECT().DownloadFile(gomock.Any()).AnyTimes().Return(nil)

	pool, _ := NewLruPool(getter, &cfg)

	pool.cfg.Cfg.MaxCacheFuncNum = 10

	for i := 0; i < 15; i++ {
		println("get:", i)
		instance := env.FunctionInstance{
			FuncID:      fmt.Sprint(i),
			StorageType: common.HDFS,
			RemotePath:  "test",
		}
		pool.getFunc(&instance)
	}

	println("====================")
	idx := 14
	for e := pool.cacheQueue.Front(); e != nil; e = e.Next() {
		assert.Equal(t, fmt.Sprint(idx), e.Value.(*env.FunctionInstance).FuncID)
		//println(e.Value.(*env.FunctionInstance).FuncID)
		idx -= 1
	}
}

func Test_getFunc_Renew(t *testing.T) {
	info := model.RuntimeInfo{
		CfgType: model.Local,
		IsDebug: true,
	}

	cfg := configs.NewConfig(info)

	ctrl := gomock.NewController(t)
	getter := mock.NewMockGetter(ctrl)
	getter.EXPECT().DownloadFile(gomock.Any()).AnyTimes().Return(nil)

	pool, _ := NewLruPool(getter, &cfg)

	pool.cfg.Cfg.MaxCacheFuncNum = 10

	for i := 0; i < 15; i++ {
		println("get:", i)
		instance := env.FunctionInstance{
			FuncID:      fmt.Sprint(i),
			StorageType: common.HDFS,
			RemotePath:  "test",
		}
		pool.getFunc(&instance)
	}

	new1 := env.FunctionInstance{
		FuncID:      "3",
		StorageType: common.HDFS,
		RemotePath:  "test",
	}
	pool.getFunc(&new1)

	new2 := env.FunctionInstance{
		FuncID:      "6",
		StorageType: common.HDFS,
		RemotePath:  "test",
	}
	pool.getFunc(&new2)

	new3 := env.FunctionInstance{
		FuncID:      "9",
		StorageType: common.HDFS,
		RemotePath:  "test",
	}
	pool.getFunc(&new3)

	println("====================")
	//idx := 14
	for e := pool.cacheQueue.Front(); e != nil; e = e.Next() {
		//assert.Equal(t, fmt.Sprint(idx), e.Value.(*env.FunctionInstance).FuncID)
		println(e.Value.(*env.FunctionInstance).FuncID)
		//idx -= 1
	}
}
