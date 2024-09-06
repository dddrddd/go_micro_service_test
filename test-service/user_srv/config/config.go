package config

type MysqlConfig struct {
	User     string `json:"user" mapstructure:"user"`
	Password string `json:"password" mapstructure:"password"`
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	Name     string `json:"name" mapstructure:"name"`
}

type ConsulConfig struct {
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`
}

type ServerConfig struct {
	Name       string       `json:"name" mapstructure:"name"`
	MysqlInfo  MysqlConfig  `json:"mysql" mapstructure:"mysql"`
	ConsulInfo ConsulConfig `json:"consul" mapstructure:"consul"`
}
