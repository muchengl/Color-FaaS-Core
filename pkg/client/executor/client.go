package executor

import (
	pb "Color-FaaS-Core/pkg/proto/executor"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type executorClient struct {
	executorIP   string
	executorPort string
	client       pb.ExecutorClient
}

func New(ip string, port string) (executorClient, error) {
	exe := executorClient{}
	exe.executorIP = ip
	exe.executorPort = port

	conn, err := grpc.Dial(exe.executorIP+":"+exe.executorPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	exe.client = pb.NewExecutorClient(conn)

	return exe, nil
}

func (e *executorClient) Heartbeat() (*pb.HeartbeatReply, error) {
	req := pb.HeartbeatRequest{
		Msg: "test",
	}
	reply, _ := e.client.Heartbeat(context.Background(), &req)
	log.Default().Printf("reply %s", reply)

	return reply, nil
}
