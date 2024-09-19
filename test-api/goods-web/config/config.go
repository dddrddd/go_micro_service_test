package config

type GoodsSrvConfig struct {
	Ip   string `mapstructure:"ip" json:"ip"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port string `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	GoodsInfo  GoodsSrvConfig `mapstructure:"goods_srv" json:"goods_srv"`
	JWT        JWTConfig      `mapstructure:"jwt" json:"JWT"`
	ConsulInfo ConsulConfig   `mapstructure:"consul" json:"consul"`
	Port       int            `mapstructure:"port" json:"port"`
}
type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	NameSpace string `mapstructure:"namespace"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}
