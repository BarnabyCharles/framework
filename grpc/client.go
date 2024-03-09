package grpc

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/BarnabyCharles/framework/consul"
)

func ClientGrpc(SeverName string) (*grpc.ClientConn, error) {
	ip, port, err2 := consul.GetClient(SeverName)
	if err2 != nil {
		return nil, err2
	}
	log.Println("client连接===============================  grpc:", ip, port)

	//ip, port, err3 := config.GetNacosServerGrpc(SeverName)
	//if err3 != nil {
	//	return nil, err3
	//}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
