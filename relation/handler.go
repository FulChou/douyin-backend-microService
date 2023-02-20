package main

import (
	"context"
	"douyin_backend_microService/pkg/errno"
	relationdemo "douyin_backend_microService/relation/kitex_gen/relationdemo"
	"douyin_backend_microService/relation/pack"
	"douyin_backend_microService/relation/rpc"
	"douyin_backend_microService/relation/service"
	"douyin_backend_microService/user/kitex_gen/userdemo"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// Relation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) Relation(ctx context.Context, req *relationdemo.RelationActionRequest) (resp *relationdemo.RelationActionResponse, err error) {
	// TODO: Your code here...
	resp = new(relationdemo.RelationActionResponse)
	relationService := service.NewRelationService(ctx)

	if req.ActionType == int64(relationdemo.RelationActionType_FollowAction) {
		err = relationService.CreateFollow(req.UserId, req.ToUserId)
	} else if req.ActionType == int64(relationdemo.RelationActionType_CancelAction) {
		err = relationService.ConcealFollow(req.UserId, req.ToUserId)
	}

	if err != nil {
		resp.Baseresp = pack.BuildResponeseMessage(err)
		return resp, err
	}

	resp.Baseresp = pack.BuildResponeseMessage(errno.Success)
	return resp, nil
}

// GetFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollow(ctx context.Context, req *relationdemo.RelationFollowListRequest) (resp *relationdemo.RelationFollowListResponse, err error) {
	// TODO: Your code here...
	resp = new(relationdemo.RelationFollowListResponse)
	relationService := service.NewRelationService(ctx)
	followList, err := relationService.GetRelationFollowList(req.UserId)
	if err != nil {
		resp.BaseResp = pack.BuildResponeseMessage(err)
		return resp, err
	}
	userIds := relationService.GetFollowUserIds(followList)
	usersmap, err := rpc.MGetUser(ctx, &userdemo.MGetUserRequest{UserIds: userIds})

	if err != nil {
		resp.BaseResp = pack.BuildResponeseMessage(err)
		return resp, err
	}

	res := make([]*userdemo.User, 0)
	for _, v := range usersmap {
		res = append(res, v)
	}

	resp.UserList = res
	resp.BaseResp = pack.BuildResponeseMessage(errno.Success)
	return resp, nil
}

// GetFollower implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollower(ctx context.Context, req *relationdemo.RelationFollowerListRequest) (resp *relationdemo.RelationFollowerListResponse, err error) {
	// TODO: Your code here...
	resp = new(relationdemo.RelationFollowerListResponse)
	relationService := service.NewRelationService(ctx)
	followList, err := relationService.GetRelationFollowList(req.UserId)
	if err != nil {
		resp.Baseresp = pack.BuildResponeseMessage(err)
		return resp, err
	}
	userIds := relationService.GetFollowerUserIds(followList)
	usersmap, err := rpc.MGetUser(ctx, &userdemo.MGetUserRequest{UserIds: userIds})

	if err != nil {
		resp.Baseresp = pack.BuildResponeseMessage(err)
		return resp, err
	}

	res := make([]*userdemo.User, 0)
	for _, v := range usersmap {
		res = append(res, v)
	}

	resp.UserList = res
	resp.Baseresp = pack.BuildResponeseMessage(errno.Success)
	return resp, nil
}

// GetFriend implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriend(ctx context.Context, req *relationdemo.RelationFriendListRequest) (resp *relationdemo.RelationFriendListResponse, err error) {
	// TODO: Your code here...
	resp = new(relationdemo.RelationFriendListResponse)
	relationService := service.NewRelationService(ctx)
	list, err := relationService.GetFriendsList(req.UserId)
	if err != nil {
		resp.Baseresp = pack.BuildResponeseMessage(err)
		return resp, err
	}
	usersmap, err := rpc.MGetUser(ctx, &userdemo.MGetUserRequest{UserIds: list})
	res := make([]*relationdemo.FriendUser, 0)
	for _, v := range usersmap {
		res = append(res, &relationdemo.FriendUser{User: v})
	}
	//todo: 获取friend最新消息

	resp.UserList = res
	resp.Baseresp = pack.BuildResponeseMessage(errno.Success)
	return resp, nil
}
