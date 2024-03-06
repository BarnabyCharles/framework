package app

import (
	"github.com/BarnabyCharles/framework/config"
	"github.com/BarnabyCharles/framework/databases/mysql"
	"github.com/BarnabyCharles/framework/databases/redis"
)

func Init(serverName string, str ...string) error {
	var err error
	err = config.ClientNacos()
	if err != nil {
		return err
	}
	//err=config.RegisterServer()
	//if err != nil {
	//	return err
	//}
	//err = consul.ConsulClient()
	//
	//if err != nil {
	//	return err
	//}
	redis.InitRedis()
	for _, val := range str {
		switch val {
		case "mysql":
			err = mysql.InitMysql(serverName)

		}
	}
	return err
}
