package rpc

import (
	"context"
	"douyin_backend_microService/pkg/constants"
	"douyin_backend_microService/pkg/errno"
	"douyin_backend_microService/pkg/mw"
	"douyin_backend_microService/user/kitex_gen/userdemo"
	"douyin_backend_microService/user/kitex_gen/userdemo/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func InitUser() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constants.ApiServiceName),
		provider.WithExportEndpoint(constants.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func CreateUser(ctx context.Context, request *userdemo.CreateUserRequest) (int64, error) {
	r, err := userClient.CreateUser(ctx, request)
	if err != nil {
		return 0, err
	}
	return r.UserId, nil

}

func CheckUser(ctx context.Context, request *userdemo.CheckUserRequest) (int64, error) {
	resp, err := userClient.CheckUser(ctx, request)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserId, nil
}
