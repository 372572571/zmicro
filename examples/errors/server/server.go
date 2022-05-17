package main

import (
	"context"
	"github.com/iobrother/zmicro"
	"github.com/iobrother/zmicro/core/log"
	"github.com/iobrother/zmicro/examples/errors/errno"
	"github.com/iobrother/zmicro/examples/proto"
	"github.com/smallnest/rpcx/server"
)

func main() {
	app := zmicro.New(zmicro.InitRpcServer(InitRpcServer))

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func InitRpcServer(s *server.Server) error {
	if err := s.RegisterName("Greeter", &GreeterImpl{}, ""); err != nil {
		return err
	}
	return nil
}

type GreeterImpl struct{}

func (s *GreeterImpl) SayHello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloReply) (err error) {
	return errno.ErrBizError()
}
