package manager

import (
	"Color-FaaS-Core/pkg/common"
	"Color-FaaS-Core/pkg/configs"
	"Color-FaaS-Core/pkg/executor/env"
	"container/list"
	"log"
	"sync"
)

type lruPool struct {
	getter env.Getter
	cfg    *configs.ExecutorConfig

	mu           sync.Mutex
	cacheQueue   *list.List
	cacheFuncMap map[string]*list.Element // funcId,idx
}

func NewLruPool(getter env.Getter, cfg *configs.ExecutorConfig) (*lruPool, error) {
	p := lruPool{
		getter:       getter,
		cfg:          cfg,
		cacheQueue:   list.New(),
		cacheFuncMap: map[string]*list.Element{},
	}
	log.Printf("init lru pool for executor, maxCacheFuncNum, %v", p.getMaxCacheLen())
	return &p, nil
}

func (l *lruPool) getMaxCacheLen() int {
	return l.cfg.Cfg.MaxCacheFuncNum
}

func (l *lruPool) removeTail() {
	tail := l.cacheQueue.Back()
	v := l.cacheQueue.Remove(tail).(*env.FunctionInstance)
	delete(l.cacheFuncMap, v.FuncID)
}

func (l *lruPool) getFunc(instance *env.FunctionInstance) (*env.FunctionInstance, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// func already exist
	if v, ok := l.cacheFuncMap[instance.FuncID]; ok {
		l.cacheQueue.Remove(v)
		l.cacheQueue.PushFront(v.Value.(*env.FunctionInstance))

		in := v.Value.(*env.FunctionInstance)
		//l.cacheFuncMap[in.FuncID]=

		return in, nil
	}

	// download this func
	err := l.downloadFunc(instance)
	if err != nil {
		return nil, err
	}

	// add to queue and map
	l.cacheQueue.PushFront(instance)
	l.cacheFuncMap[instance.FuncID] = l.cacheQueue.Front()

	if l.cacheQueue.Len() > l.getMaxCacheLen() {
		l.removeTail()
	}

	return instance, nil
}

func (l *lruPool) downloadFunc(instance *env.FunctionInstance) error {
	// download func file
	err := l.getter.DownloadFile(*instance)
	if err != nil {
		log.Printf("download function fail: %v", err)
		return err
	}
	instance.Status = common.Init
	return nil
}
