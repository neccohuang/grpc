package main

import (
	"context"
	"log"
	"time"

	pb "github.com/neccohuang/grpc/proto"
)

type EchoServer struct{}

func (e *EchoServer) Echo(ctx context.Context, req *pb.EchoRequest) (resp *pb.EchoReply, err error) {

	log.Printf("receive client request, client send:%s\n", req.Msg)
	return &pb.EchoReply{
		Msg:      req.Msg,
		Unixtime: time.Now().Unix(),
	}, nil

}

func main() {

}
