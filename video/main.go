package main

import (
	"douyin_backend_microService/api/biz/mw"
	"douyin_backend_microService/pkg/constants"
	"douyin_backend_microService/video/dal/db"
	"douyin_backend_microService/video/kitex_gen/videodemo/videoservice"
	"douyin_backend_microService/video/rpc"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	registry, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})

	addr, err := net.ResolveTCPAddr(constants.TCP, constants.VideoServiceAddr)
	if err != nil {
		panic(err)
	}
	svr := videoservice.NewServer(new(VideoServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(registry),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.VideoServiceName}))
	if err != nil {
		panic(err)
	}

	mw.InitJWT()
	rpc.InitUser()
	db.Init()

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
