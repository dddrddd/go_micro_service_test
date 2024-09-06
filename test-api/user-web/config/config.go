package config

type UserSrvConfig struct {
	Ip   string `mapstructure:"ip"`
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	UserInfo   UserSrvConfig `mapstructure:"user_srv"`
	JWT        JWTConfig     `mapstructure:"jwt"`
	Redis      RedisConfig   `mapstructure:"redis"`
	ConsulInfo ConsulConfig  `mapstructure:"consul"`
	Port       int           `mapstructure:"port"`
}
