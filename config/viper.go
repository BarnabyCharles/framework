package config

import (
	"errors"

	"github.com/spf13/viper"
)

var Nacoscfg *AppConfig

func InitViper(fileName, filePath string) (string, int, string, string, error) {
	v := viper.New()

	//v.SetConfigName("Nao")
	v.SetConfigFile(filePath)

	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		return "", 0, "", "", errors.New("获取配置信息失败！" + err.Error())
	}
	err = v.Unmarshal(&Nacoscfg)

	if err != nil {
		return "", 0, "", "", errors.New("获取配置文件nacos配置信息失败！" + err.Error())
	}
	host := Nacoscfg.Nacos.Host
	port := Nacoscfg.Nacos.Port
	serverName := Nacoscfg.Nacos.ServerName
	NamespaceId := Nacoscfg.Nacos.NamespaceId
	return host, port, serverName, NamespaceId, nil
}
func InitViper2(fileName, filePath string) (*viper.Viper, error) {
	v := viper.New()

	//v.SetConfigName(fileName)
	v.SetConfigFile(filePath)

	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, errors.New("读取配置信息失败！" + err.Error())
	}
	return v, nil
}
