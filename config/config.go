package config

type AppConfig struct {
	App   string      `yaml:"App"`
	Mysql MysqlConfig `yaml:"Mysql"`
}

type MysqlConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}
