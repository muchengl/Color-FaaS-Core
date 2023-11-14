package funcmanager

import (
	"Color-FaaS-Core/pkg/common"
	"log"
)

type lruPool struct {
	getter *funcGetter
}

func NewLruPool(getter *funcGetter) (*lruPool, error) {
	p := lruPool{
		getter: getter,
	}
	return &p, nil
}

func (l *lruPool) getFunc(instance *FunctionInstance) (*FunctionInstance, error) {

	// todo lru pool, now there is no pool
	l.downloadFunc(instance)

	return instance, nil
}

func (l *lruPool) downloadFunc(instance *FunctionInstance) error {
	// download func file
	err := l.getter.downloadFile(*instance)
	if err != nil {
		log.Printf("download function fail: %v", err)
		return err
	}
	instance.Status = common.Init
	return nil
}
