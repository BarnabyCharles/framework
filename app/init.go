package app

import "github.com/BarnabyCharles/framework/mysql"

func Init(str ...string) error {
	var err error
	for _, val := range str {
		switch val {
		case "mysql":
			err = mysql.InitMysql()

		}
	}
	return err
}
