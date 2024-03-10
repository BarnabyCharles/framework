package app

import (
	"github.com/BarnabyCharles/framework/consul"

	"github.com/BarnabyCharles/framework/config"
	"github.com/BarnabyCharles/framework/databases/mysql"
)

func Init(fileName, filePath string, str ...string) error {
	var err error
	nacoscfg, err := config.InitViper(fileName, filePath)
	if err != nil {
		return err
	}
	//host, port, serverName, NamespaceId
	host := nacoscfg.Nacos.Host
	port := nacoscfg.Nacos.Port
	serverName := nacoscfg.Nacos.ServerName
	NamespaceId := nacoscfg.Nacos.NamespaceId
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
