package main

import (
	"douyin_backend_microService/pkg/constants"
	"douyin_backend_microService/relation/dal/db"
	relationdemo "douyin_backend_microService/relation/kitex_gen/relationdemo/relationservice"
	"douyin_backend_microService/relation/rpc"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	registry, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})

	addr, err := net.ResolveTCPAddr(constants.TCP, constants.RelationServiceAddr)
	if err != nil {
		panic(err)
	}
	svr := relationdemo.NewServer(new(RelationServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(registry),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.RelationServiceName}))
	if err != nil {
		panic(err)
	}
	rpc.InitUser()
	db.Init()
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
