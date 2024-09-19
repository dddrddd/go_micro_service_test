package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type Registry struct {
	Host string
	Port int
}

type RegistryClient interface {
	Register(address string, port int, name string, tags []string, id string) error
	Deregister(id string) error
}

func NewRegistryClient(host string, port int) RegistryClient {
	return &Registry{
		Host: host,
		Port: port,
	}
}

func (r *Registry) Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//开始注册
	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.ID = id
	registration.Name = name
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	//配置检查项
	check := new(api.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d/health", address, port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "10s"
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

func (r *Registry) Deregister(id string) error {
	return nil
}
