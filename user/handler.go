package main

import (
	"context"
	"douyin_backend_microService/pkg/errno"
	"douyin_backend_microService/user/kitex_gen/userdemo"
	"douyin_backend_microService/user/pack"
	"douyin_backend_microService/user/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	ctx context.Context
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *userdemo.GetUserRequest) (resp *userdemo.GetUserResponse, err error) {
	// TODO: Your code here...
	userService := service.NewUserService(ctx)
	resp = new(userdemo.GetUserResponse)
	user, err := userService.GetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildResponeseMessage(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildResponeseMessage(errno.Success)
	resp.User = pack.ConcertUser(user)
	return resp, nil
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *userdemo.MGetUserRequest) (resp *userdemo.MGetUserResponse, err error) {
	// TODO: Your code here...
	userService := service.NewUserService(ctx)
	resp = new(userdemo.MGetUserResponse)
	users, err := userService.MGetUser(req)

	if err != nil {
		resp.BaseResp = pack.BuildResponeseMessage(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildResponeseMessage(errno.Success)
	resp.Users = users

	return resp, nil
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *userdemo.CreateUserRequest) (resp *userdemo.CreateUserResponse, err error) {
	// TODO: Your code here...
	resp = new(userdemo.CreateUserResponse)
	userService := service.NewUserService(ctx)
	userId, err := userService.CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildResponeseMessage(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildResponeseMessage(errno.Success)
	resp.UserId = userId

	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userdemo.CheckUserRequest) (resp *userdemo.CheckUserResponse, err error) {
	// TODO: Your code here...
	resp = new(userdemo.CheckUserResponse)
	userService := service.NewUserService(ctx)
	userId, err := userService.CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildResponeseMessage(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildResponeseMessage(err)
	resp.UserId = userId
	return resp, nil
}
