package utils

import (
	consulapi "github.com/hashicorp/consul/api"
	"log"
)

var ConsulClient *consulapi.Client
var ServiceId string
var ServiceName string
var ServicePort int
func init()  {
	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	ConsulClient = client
}

func RegService() {
	//设置注册信息
	reg := consulapi.AgentServiceRegistration{}
	reg.ID = "gameService"
	reg.Name = "gameService"
	reg.Address = "127.0.0.1"
	reg.Port = 8000
	reg.Tags = []string{"primary"}

	//设置健康检查
	check := consulapi.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP = "http://127.0.0.1:8000/health"
	//check.HTTP = fmt.Sprintf("http://%s:%d/health",reg.Address,ServicePort)
	reg.Check = &check

	//注册
	err := ConsulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal(err)
	}
}

//反注册
func Unregservice()  {
	ConsulClient.Agent().ServiceDeregister("gameService")
}