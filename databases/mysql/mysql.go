package mysql

import (
	"errors"
	"fmt"
	"log"

	"github.com/ghodss/yaml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/BarnabyCharles/framework/config"
)

var DB *gorm.DB

func InitMysql(serverName string) error {

	nacos, err := config.GetNacosConfig(serverName, "DEFAULT_GROUP")
	if err != nil {
		return err
	}
	var mysqConfig config.AppConfig
	err = yaml.Unmarshal([]byte(nacos), &mysqConfig)

	if err != nil {
		return errors.New("将yaml文件转换为结构体格式失败！" + err.Error())
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		mysqConfig.Mysql.Username,
		mysqConfig.Mysql.Password,
		mysqConfig.Mysql.Host,
		mysqConfig.Mysql.Port,
		mysqConfig.Mysql.Database,
	)
	log.Println("数据库连接=====================", dsn)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return errors.New("连接数据库失败！" + err.Error())
	}

	DB = db

	return err
}

func WithTx(txFc func(tx *gorm.DB) error) {
	var err error
	tx := DB.Begin()
	err = txFc(tx)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
}
