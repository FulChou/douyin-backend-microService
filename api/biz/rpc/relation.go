package rpc

import (
	"context"
	"douyin_backend_microService/pkg/constants"
	"douyin_backend_microService/pkg/errno"
	"douyin_backend_microService/pkg/mw"
	"douyin_backend_microService/relation/kitex_gen/relationdemo"
	"douyin_backend_microService/relation/kitex_gen/relationdemo/relationservice"
	"douyin_backend_microService/user/kitex_gen/userdemo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var relationClient relationservice.Client

func InitRelation() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constants.ApiServiceName),
		provider.WithExportEndpoint(constants.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := relationservice.NewClient(
		constants.RelationServiceName,
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
	relationClient = c
}

func Relation(ctx context.Context, req *relationdemo.RelationActionRequest) error {
	resp, err := relationClient.Relation(ctx, req)
	if err != nil {
		klog.Error("relation err, ", err.Error())
		return err
	}
	if resp.Baseresp.StatusCode != 0 {
		klog.Error("relation err, ", resp.Baseresp.StatusMsg)
		return errno.NewErrNo(resp.Baseresp.StatusCode, resp.Baseresp.StatusMsg)
	}

	return nil

}

func GetFollowList(ctx context.Context, req *relationdemo.RelationFollowListRequest) ([]*userdemo.User, error) {
	resp, err := relationClient.GetFollow(ctx, req)
	if err != nil {
		klog.Error("get follow list err, ", err.Error())
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		klog.Error("get follow err, ", resp.BaseResp.StatusMsg)
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}

	return resp.UserList, nil
}

func GetFollowerList(ctx context.Context, req *relationdemo.RelationFollowerListRequest) ([]*userdemo.User, error) {
	resp, err := relationClient.GetFollower(ctx, req)
	if err != nil {
		klog.Error("get follower err, ", err.Error())
		return nil, err
	}
	if resp.Baseresp.StatusCode != 0 {
		klog.Error("get follower err, ", resp.Baseresp.StatusMsg)
		return nil, errno.NewErrNo(resp.Baseresp.StatusCode, resp.Baseresp.StatusMsg)
	}

	return resp.UserList, nil

}

func GetFriendList(ctx context.Context, req *relationdemo.RelationFriendListRequest) ([]*relationdemo.FriendUser, error) {
	resp, err := relationClient.GetFriend(ctx, req)
	if err != nil {
		klog.Error("get friend err, ", err.Error())
		return nil, err
	}
	if resp.Baseresp.StatusCode != 0 {
		klog.Error("get friend err, ", resp.Baseresp.StatusMsg)
		return nil, errno.NewErrNo(resp.Baseresp.StatusCode, resp.Baseresp.StatusMsg)
	}

	return resp.UserList, nil

}
