package main

import (
	"douyin_backend_microService/message/dal/db"
	messagedemo "douyin_backend_microService/message/kitex_gen/messagedemo/messageservice"
	"douyin_backend_microService/pkg/constants"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	registry, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})

	addr, err := net.ResolveTCPAddr(constants.TCP, constants.MessageServiceAddr)
	if err != nil {
		panic(err)
	}

	svr := messagedemo.NewServer(new(MessageServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(registry),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.MessageServiceName}))

	db.Init()

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}
