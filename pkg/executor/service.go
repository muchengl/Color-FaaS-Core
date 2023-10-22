package executor

import (
	"Color-FaaS-Core/pkg/common"
	fmgr "Color-FaaS-Core/pkg/executor/funcmanager"
	model "Color-FaaS-Core/pkg/model"
	pb "Color-FaaS-Core/pkg/proto/executor"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
)

type Executor struct {
	pb.UnimplementedExecutorServer

	executorID  string
	RuntimeInfo model.RuntimeInfo
	funcManager fmgr.Service
	cfg         config
}

func New(info model.RuntimeInfo) (*Executor, error) {
	executor := Executor{}
	executor.RuntimeInfo = info
	executor.cfg = newConfig(info)
	executor.executorID = uuid.New().String()

	mgr, err := fmgr.New(info)
	if err != nil {
		log.Fatalf("Init executor err %v", err)
		return nil, err
	}
	executor.funcManager = mgr

	return &executor, nil
}

func (exe *Executor) InitRunningEnv() error {
	err := os.Mkdir(exe.cfg.Cfg.FuncRunDir, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Mkdir(exe.cfg.Cfg.FuncFilePath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (exe *Executor) Start() {
	lis, err := net.Listen("tcp", "127.0.0.1:"+exe.cfg.Cfg.Port)
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

func (e *Executor) InitTask(ctx context.Context, req *pb.TaskInstance) (*pb.InitTaskReply, error) {
	log.Default().Print("InitTask access")

	reply := pb.InitTaskReply{}

	funcInstance := fmgr.FunctionInstance{
		FuncName:    req.FuncName,
		FuncID:      req.FuncID,
		StorageType: req.FuncStorageType,
		RemotePath:  req.TaskFuncPath,
		LocalPath:   e.cfg.Cfg.FuncFilePath + "/" + uuid.NewString(),
		Status:      common.NotInit,
	}
	err := e.funcManager.Init(funcInstance)
	if err != nil {
		log.Printf("init func fail, %v", funcInstance)
		reply.Status = -1
		reply.Msg = "init func fail"
		return &reply, nil
	}

	err = e.funcManager.Start()
	if err != nil {
		log.Printf("start func fail, %v", funcInstance)
		reply.Status = -1
		reply.Msg = "start func fail"
		return &reply, nil
	}

	reply.Msg = "success"
	reply.Status = 1
	return &reply, nil
}

func (e *Executor) RunTask(ctx context.Context, req *pb.RunTaskRequest) (*pb.RunTaskReply, error) {

	return nil, status.Errorf(codes.Unimplemented, "method RunTask not implemented")
}

func (e *Executor) KillTask(ctx context.Context, req *pb.KillTaskRequest) (*pb.KillTaskReply, error) {
	reply := pb.KillTaskReply{}
	err := e.funcManager.Kill()
	if err != nil {
		reply.Status = -1
		reply.Msg = "kill func fail"
		return &reply, nil
	}

	reply.Msg = "success"
	reply.Status = 1
	return &reply, nil
}
