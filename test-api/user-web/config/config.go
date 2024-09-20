package config

type UserSrvConfig struct {
	Ip   string `mapstructure:"ip" json:"ip"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type RedisConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	Name       string        `mapstructure:"name" json:"name"`
	Host       string        `mapstructure:"host" json:"host"`
	Tags       []string      `mapstructure:"tags" json:"tags"`
	UserInfo   UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	JWT        JWTConfig     `mapstructure:"jwt" json:"JWT"`
	Redis      RedisConfig   `mapstructure:"redis" json:"redis"`
	ConsulInfo ConsulConfig  `mapstructure:"consul" json:"consul"`
	Port       int           `mapstructure:"port" json:"port"`
}
type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	NameSpace string `mapstructure:"namespace"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}
