package grpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ghodss/yaml"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/BarnabyCharles/framework/config"
)

func RegisterGRPC(serverName string, register func(s *grpc.Server), cert, key string) error {

	nacosConfig, err := config.GetNacosConfig(serverName, "DEFAULT_GROUP")
	if err != nil {
		return err
	}
	var AppConfig config.AppConfig
	err = yaml.Unmarshal([]byte(nacosConfig), &AppConfig)
	if err != nil {
		return err
	}

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", AppConfig.Host, AppConfig.Port))
	if err != nil {
		log.Panicf("failed to listen%v", err)
		return err
	}

	s := grpc.NewServer()
	// 反射查询
	reflection.Register(s)
	register(s)

	//err = consul.AgentService(AppConfig.Host, AppConfig.Port)
	//if err != nil {
	//	return err
	//}
	//
	//// 注册健康检测
	//grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	log.Println("server", listen.Addr())
	go func() {
		if err := s.Serve(listen); err != nil {
			log.Panicf("failed to server%v", err)

		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	_, err = config.ClientServer.DeregisterInstance(config.Deregister())
	if err != nil {
		log.Println("注销服务实例失败！", err)
		return err
	}
	//s.GracefulStop()

	return err
}
