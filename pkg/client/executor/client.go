package executor

import (
	pb "Color-FaaS-Core/pkg/proto/executor"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Client struct {
	executorIP   string
	executorPort string
	client       pb.ExecutorClient
}

func New(ip string, port string) (Client, error) {
	exe := Client{}
	exe.executorIP = ip
	exe.executorPort = port

	conn, err := grpc.Dial(exe.executorIP+":"+exe.executorPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	exe.client = pb.NewExecutorClient(conn)

	return exe, nil
}

func (e *Client) Heartbeat() (*pb.HeartbeatReply, error) {
	req := pb.HeartbeatRequest{
		Msg: "test",
	}
	reply, _ := e.client.Heartbeat(context.Background(), &req)
	log.Default().Printf("reply %s", reply)

	return reply, nil
}

func (e *Client) InitTask(instance *pb.TaskInstance) (*pb.InitTaskReply, error) {
	reply, _ := e.client.InitTask(context.Background(), instance)
	log.Default().Printf("reply %s", reply)

	return reply, nil
}

func (e *Client) RunTask(instance *pb.TaskInstance) (*pb.RunTaskReply, error) {
	reply, _ := e.client.RunTask(context.Background(), instance)
	log.Default().Printf("reply %s", reply)

	return reply, nil
}

func (e *Client) KillTask(instance *pb.TaskInstance) (*pb.KillTaskReply, error) {
	reply, _ := e.client.KillTask(context.Background(), instance)
	log.Default().Printf("reply %s", reply)

	return reply, nil
}
