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

func initUser() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constants.VideoServiceName),
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

func MGetUser(ctx context.Context, req *userdemo.MGetUserRequest) (map[int64]*userdemo.User, error) {
	r, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if r.BaseResp.StatusCode != 0 {
		return nil, errno.ErrNo{
			ErrCode: r.BaseResp.StatusCode,
			ErrMsg:  r.BaseResp.StatusMsg,
		}
	}

	users := make(map[int64]*userdemo.User, 0)
	for _, u := range r.Users {
		users[u.Id] = u
	}

	return users, nil
}

func GetUser(ctx context.Context, req *userdemo.GetUserRequest) (*userdemo.User, error) {
	r, err := userClient.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if r.BaseResp.StatusCode != 0 {
		return nil, errno.ErrNo{
			ErrCode: r.BaseResp.StatusCode,
			ErrMsg:  r.BaseResp.StatusMsg,
		}
	}
	return r.User, nil

}
