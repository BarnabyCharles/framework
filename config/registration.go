package config

import (
	"log"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var ClientServer naming_client.INamingClient

func RegisterServer() error {
	// 配置Nacos服务器地址和命名空间等信息
	sc := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}

	cc := &constant.ClientConfig{
		NamespaceId:         "", //
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogLevel:            "debug",
	}

	// 创建nacos客户端
	var err error
	ClientServer, err = clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig:  cc,
		ServerConfigs: sc,
	})
	if err != nil {
		log.Println("连接nacos配置客户端失败！", err)
	}
	suesscc, err := ClientServer.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "10.2.171.84",
		Port:        8081,
		Weight:      0,
		Enable:      true,
		Healthy:     true, // 开启健康检测
		Metadata:    nil,
		ClusterName: "",
		ServiceName: "user_srv",
		GroupName:   "DEFAULT_GROUP",
		Ephemeral:   false,
	})
	if err != nil {
		log.Println("注册服务出错！", err)
		return err
	}
	if suesscc {
		log.Println("nacos服务注册成功！", err)
		return err
	} else {
		log.Println("nacos服务注册失败！", err)
		return err
	}

}

func GetNacosServerGrpc(serviceName string) (string, int, error) {

	GetServiceParam := vo.GetServiceParam{
		Clusters:    nil,
		ServiceName: serviceName,
		GroupName:   "DEFAULT_GROUP",
	}

	service, err := ClientServer.GetService(GetServiceParam)
	if err != nil {
		log.Println("获取nacos服务发现失败！", err)
		return "", 0, nil
	}
	var host string
	var port int
	for _, val := range service.Hosts {
		host = val.Ip
		port = int(val.Port)
	}

	log.Println("dsjklsdfjjfdskjkldsf;dsfjkl;", host, port)

	return host, port, nil
}

func Deregister() vo.DeregisterInstanceParam {
	instance := vo.DeregisterInstanceParam{
		Ip:          "10.2.171.84",
		Port:        8084,
		ServiceName: "user_srv",
		//Cluster:     "your_service_cluster",
	}
	return instance
}
