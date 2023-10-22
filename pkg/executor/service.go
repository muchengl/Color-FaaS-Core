package executor

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	mgr "Color-FaaS-Core/pkg/executor/fmanager"
	model "Color-FaaS-Core/pkg/model"
	pb "Color-FaaS-Core/pkg/proto/executor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 50001, "executor port")
)

type Executor struct {
	pb.UnimplementedExecutorServer

	RuntimeInfo model.RuntimeInfo
	fmanager    mgr.FuncManager
	cfg         config
}

func New() (Executor, error) {
	executor := Executor{}
	return executor, nil
}

func (exe *Executor) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterExecutorServer(s, exe)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (e *Executor) Heartbeat(ctx context.Context, req *pb.HeartbeatRequest) (*pb.HeartbeatReply, error) {
	log.Default().Print("heartbeat access")
	reply := pb.HeartbeatReply{}
	reply.Status = 1
	reply.Msg = "alive"
	return &reply, nil
}

func (e *Executor) InitTask(ctx context.Context, req *pb.InitTaskRequest) (*pb.InitTaskReply, error) {
	log.Default().Print("InitTask access")

	return nil, status.Errorf(codes.Unimplemented, "method InitTask not implemented")
}

func (e *Executor) RunTask(ctx context.Context, req *pb.RunTaskRequest) (*pb.RunTaskReply, error) {

	return nil, status.Errorf(codes.Unimplemented, "method RunTask not implemented")
}

func (e *Executor) KillTask(ctx context.Context, req *pb.KillTaskRequest) (*pb.KillTaskReply, error) {

	return nil, status.Errorf(codes.Unimplemented, "method KillTask not implemented")
}
