package pool

import (
	config "Color-FaaS-Core/pkg/configs"
	"log"
	"os/exec"
	"time"
)
import executor "Color-FaaS-Core/pkg/client/executor"

type LruPool struct {
	cfg config.ClientConfig
}

func NewLruPool(cfg config.ClientConfig) *LruPool {
	pool := LruPool{
		cfg: cfg,
	}

	return &pool
}

// todo can't new executor
func (p *LruPool) GetExecutor() (*executor.Client, *exec.Cmd, error) {
	// start a new executor
	log.Print("start executor...")
	path := p.cfg.Cfg.ExecutorPath
	cmd := exec.Command(path)
	cmd.Dir = "../executor"
	err := cmd.Start()
	time.Sleep(time.Millisecond * 100)

	if err != nil {
		log.Printf("can run: " + path)
		return nil, cmd, err
	}

	// connect
	log.Print("connect executor...")
	exe, _ := executor.New("127.0.0.1", "50001")

	// return client
	return &exe, cmd, nil
}
