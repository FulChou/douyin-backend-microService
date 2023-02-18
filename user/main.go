package main

import (
	"douyin_backend_microService/pkg/constants"
	"douyin_backend_microService/user/dal/db"
	"douyin_backend_microService/user/kitex_gen/userdemo/userservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	registry, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})

	addr, err := net.ResolveTCPAddr(constants.TCP, constants.UserServiceAddr)
	if err != nil {
		panic(err)
	}
	svr := userservice.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(registry),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}))
	if err != nil {
		panic(err)
	}

	db.Init()
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
