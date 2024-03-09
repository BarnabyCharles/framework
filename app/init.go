package app

import (
	"github.com/BarnabyCharles/framework/consul"

	"github.com/BarnabyCharles/framework/config"
	"github.com/BarnabyCharles/framework/databases/mysql"
)

func Init(fileName, filePath string, str ...string) error {
	var err error
	host, port, serverName, NamespaceId, err := config.InitViper(fileName, filePath)
	if err != nil {
		return err
	}

	err = config.ClientNacos(NamespaceId, host, port)
	if err != nil {
		return err
	}

	//err=config.RegisterServer()
	//if err != nil {
	//	return err
	//}

	err = consul.ConsulClient(serverName)

	if err != nil {
		return err
	}
	for _, val := range str {
		switch val {
		case "mysql":
			err = mysql.InitMysql(serverName)

		}
	}
	return err
}
