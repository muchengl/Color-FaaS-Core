package clientzk

import (
	config "Color-FaaS-Core/pkg/configs"
	"fmt"
	zk "github.com/go-zookeeper/zk"
	"time"
)

type ZKService struct {
	conn *zk.Conn
	cfg  config.ClientConfig
}

func NewZKService(cfg config.ClientConfig) (*ZKService, error) {
	zks := ZKService{
		cfg: cfg,
	}
	conn, _, err := zk.Connect([]string{cfg.Cfg.ZKPath}, time.Second)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ZooKeeper: %v", err)
	}
	zks.conn = conn

	return &zks, nil
}

func (z *ZKService) RegisterService(path string, data []byte) error {
	flags := int32(zk.FlagEphemeral)
	acl := zk.WorldACL(zk.PermAll)

	_, err := z.conn.Create(path, data, flags, acl)
	if err != nil {
		return fmt.Errorf("failed to create a node: %v", err)
	}

	fmt.Println("Service registered successfully")
	return nil
}
