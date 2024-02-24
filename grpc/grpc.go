package grpc

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RegisterGRPC(port int, register func(s *grpc.Server)) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Panicf("failed to listen%v", err)
		return err
	}
	server := grpc.NewServer()
	// 反射查询
	reflection.Register(server)
	register(server)
	err = server.Serve(listen)
	if err != nil {
		log.Panicf("failed to server%v", err)
		return err
	}
	return err
}
