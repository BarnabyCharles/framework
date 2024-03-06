package consul

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
)

var ConsulCli *api.Client
var Srvid string

func ConsulClient() error {
	var err error
	ConsulCli, err = api.NewClient(api.DefaultConfig())
	if err != nil {
		return errors.New("连接consul客户端失败！" + err.Error())
	}
	return nil
}

func AgentService(Address string, Port int) error {
	Srvid = uuid.New().String()

	check := &api.AgentServiceCheck{
		Interval:                       "5s",
		Timeout:                        "5s",
		GRPC:                           fmt.Sprintf("%s:%d", Address, Port),
		DeregisterCriticalServiceAfter: "30s",
	}
	err := ConsulCli.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      Srvid,
		Name:    "user_srv",
		Tags:    []string{"GRPC"},
		Port:    Port,
		Address: Address,
		Check:   check,
	})
	if err != nil {
		return err
	}
	return nil
}

func GetClient(serverName string) (string, int, error) {
	name, data, err := ConsulCli.Agent().AgentHealthServiceByName(serverName)
	if name != "passing" {
		log.Println("获取consul服务发现失败！", err)
		return "", 0, nil
	}
	var Address string
	var Port int
	for _, val := range data {
		Address = val.Service.Address
		Port = val.Service.Port
	}
	log.Println("consul当中获取的服务发现端口", Address, Port)
	return Address, Port, nil
}
