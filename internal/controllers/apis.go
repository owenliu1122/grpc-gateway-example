package controllers

import (
	"context"
	"fmt"

	"github.com/owenliu1122/grpc-gateway-example/pb"
)

// APIOptions API Server 配置参数
type APIOptions struct {
	Port int `mapstructure:"port"`
}

// APIServer gRPC API Server。实现 pb.HelloWorldServer 接口，用于处理 gRPC 请求。
type APIServer struct{}

// NewAPIServer APIServer 构造函数。
func NewAPIServer(opts APIOptions) *APIServer {
	return &APIServer{}
}

func (ctl *APIServer) HelloWorld(ctx context.Context, req *pb.HelloWorldReq) (*pb.HelloWorldResp, error) {
	return &pb.HelloWorldResp{
		Echo: fmt.Sprintf("[%s]Hello World!", req.Name),
	}, nil
}
