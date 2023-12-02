package client

import (
	clientmodel "Color-FaaS-Core/pkg/client/model"
	pool "Color-FaaS-Core/pkg/client/pool"
	"Color-FaaS-Core/pkg/client/zk"
	"Color-FaaS-Core/pkg/common"
	config "Color-FaaS-Core/pkg/configs"
	model "Color-FaaS-Core/pkg/model"
	epb "Color-FaaS-Core/pkg/proto/executor"
	pb "Color-FaaS-Core/pkg/proto/scheduler"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	codeFunctionPanic    = 512
	codeInitializerPanic = 513
	functionSuccess      = 200
	functionFail         = 500
	unknownError         = 590
)

type Client struct {
	cfg         config.ClientConfig
	pool        pool.LruPool
	server      server.Hertz
	RuntimeInfo model.RuntimeInfo
	zkService   clientzk.ZKService
}

func NewClient(info model.RuntimeInfo) (*Client, error) {
	client := Client{
		cfg: config.NewClientConfig(info),
	}
	client.pool = *pool.NewLruPool(client.cfg)
	zkService, _ := clientzk.NewZKService(client.cfg)
	client.zkService = *zkService

	serviceData := []byte("127.0.0.1:9090")
	client.zkService.RegisterService("/clients", serviceData)

	return &client, nil
}

func (c *Client) Run() {
	h := server.Default(server.WithHostPorts(":" + c.cfg.Cfg.Port))

	c.server = *h

	h.GET("/run", c.invokeFunc)
	h.POST("/run", c.invokeFunc)

	h.Spin()
}

func (c *Client) invokeFunc(ctx context.Context, req *app.RequestContext) {
	reqc := clientmodel.TaskInstance{}
	if err := req.Bind(&reqc); err != nil {
		req.String(consts.StatusOK, string(codeInitializerPanic))
	}

	reqByte, _ := json.Marshal(reqc)
	log.Default().Printf("user msg: %s", string(reqByte))

	// Run it!
	out, logs, _ := c.RunFunction(reqc)

	response := pb.RunTaskReply{
		Status: functionSuccess,
		Msg:    out,
		Logs:   logs,
	}
	responseByte, _ := json.Marshal(response)
	req.String(consts.StatusOK, string(responseByte))
}

func getTaskExecutorInstance(fun clientmodel.TaskInstance) epb.TaskInstance {
	t := epb.TaskInstance{
		TaskID:          fun.TaskID,
		FuncName:        fun.FuncName,
		FuncID:          fun.FuncName,
		FuncStorageType: common.HDFS,
		TaskFuncPath:    fun.TaskFuncPath,
	}
	return t
}

func (c *Client) RunFunction(fun clientmodel.TaskInstance) (string, string, error) {
	exe, cmd, _ := c.pool.GetExecutor()
	defer SendKillRequest()
	defer cmd.Process.Kill()

	function := getTaskExecutorInstance(fun)

	var logs = ""

	// init
	replyInit, _ := exe.InitTask(&function)
	reqByte, _ := json.Marshal(replyInit)
	println(string(reqByte))

	logs += string(reqByte)
	logs += "\n"

	// start func
	replyRun, _ := exe.RunTask(&function)
	reqByte, _ = json.Marshal(replyRun)
	println(string(reqByte))

	logs += string(reqByte)
	logs += "\n"

	// send a request to func
	funcReturn, _ := SendMsgRequest(fun.TaskInput)

	return funcReturn, logs, nil
}

func SendMsgRequest(umsg string) (string, error) {
	baseURL := "http://localhost:8888/invoke?umsg=" + umsg
	queryParams := url.Values{}
	//queryParams.Set("umsg", umsg)
	urlWithParams := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())

	resp, err := http.Get(urlWithParams)
	if err != nil {
		log.Print("send req to func fail")
		return fmt.Sprintf("error sending request: %v", err), fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("received non-OK response status: %s", resp.Status), fmt.Errorf("received non-OK response status: %s", resp.Status)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Request sent successfully")
	return string(bodyBytes), nil
}

func SendKillRequest() (string, error) {
	baseURL := "http://localhost:8888/quit"
	queryParams := url.Values{}
	//queryParams.Set("umsg", umsg)
	urlWithParams := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())

	resp, err := http.Get(urlWithParams)
	if err != nil {
		log.Print("send req to func fail")
		return fmt.Sprintf("error sending request: %v", err), fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("received non-OK response status: %s", resp.Status), fmt.Errorf("received non-OK response status: %s", resp.Status)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Request sent successfully")
	return string(bodyBytes), nil
}
