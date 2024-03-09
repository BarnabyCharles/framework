package consul

import (
	"errors"
	"fmt"
	"log"

	"github.com/ghodss/yaml"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"

	"github.com/BarnabyCharles/framework/config"
)

var ConsulCli *api.Client
var Srvid string

func ConsulClient(serverName string) error {
	var consulCfg config.AppConfig
	nacosConfig, err := config.GetNacosConfig(serverName, "DEFAULT_GROUP")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(nacosConfig), &consulCfg)
	if err != nil {
		return errors.New("consul配置信息转换为结构体失败！" + err.Error())
	}
	consulConfig := api.DefaultConfig()

	ip := consulCfg.Consul.Host
	port := consulCfg.Consul.Port
	consulConfig.Address = fmt.Sprintf("%s:%d", ip, port)

	ConsulCli, err = api.NewClient(consulConfig)
	if err != nil {
		return errors.New("连接consul客户端失败！" + err.Error())
	}
	return nil
}

func AgentService(serverName, Address string, Port int) error {
	Srvid = uuid.New().String()
	check := &api.AgentServiceCheck{
		Interval:                       "5s",
		Timeout:                        "5s",
		GRPC:                           fmt.Sprintf("%s:%d", Address, Port),
		DeregisterCriticalServiceAfter: "10s",
	}
	err := ConsulCli.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      Srvid,
		Name:    serverName,
		Tags:    []string{"GRPC"},
		Port:    Port,
		Address: Address,
		Check:   check,
	})
	if err != nil {
		return errors.New("consul注册服务失败！" + err.Error())
	}
	return nil
}

//	func GetIndex(ctx context.Context, serverName string, duration time.Duration) (int, error) {
//		// 首先判断这个key 的值存不存在
//		exist := redis.ExpireKey(ctx, serverName, "consul:index")
//		var index int
//		// 存在获取 key 的值
//		if exist {
//			str, err := redis.GetRedisKey2(ctx, serverName, "consul:index")
//			if err != nil {
//				return 0, err
//			}
//			index, _ = strconv.Atoi(str)
//			if err != nil {
//				return 0, nil
//			}
//
//			if index >= 3 {
//				index = 0
//				err := redis.SetKey(ctx, serverName, "consul:index", index, duration)
//				if err != nil {
//					return 0, err
//				}
//			}
//			// 自增  key的值
//			err = redis.IndexAdd(ctx, serverName, "consul:index", duration)
//			if err != nil {
//				return 0, err
//			}
//		}
//		// 首次设置 key的值默认   0
//		err := redis.IndexAdd(ctx, serverName, "consul:index", duration)
//		if err != nil {
//			return 0, err
//		}
//		log.Println("设置index=========================", index)
//		return index, err
//	}
func GetClient(serverName string) (string, int, error) {
	name, data, err := ConsulCli.Agent().AgentHealthServiceByName("user.day05")
	if name != "passing" {
		log.Println("获取consul服务发现失败！", err)
		return "", 0, nil
	}
	var Address string
	var Port int
	//index, err := GetIndex(context.Background(), serverName, 0)
	//if err != nil {
	//	return "", 0, err
	//}
	//log.Println("设置index第一次值", index)
	//log.Println("index======================", index)
	//Address = data[index].Service.Address
	//Port = data[index].Service.Port
	for _, val := range data {
		Address = val.Service.Address
		Port = val.Service.Port
	}

	log.Println("consul当中获取的服务发现端口", Address, Port)
	return Address, Port, nil
}
