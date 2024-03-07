package config

type AppConfig struct {
	App string `yaml:"App"`

	Host string `json:"Host" yaml:"Host"`
	Port int    `json:"Port" yaml:"Port"`

	Mysql MysqlConfig `yaml:"Mysql"`
}

type MysqlConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

type RedisConfig struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}
