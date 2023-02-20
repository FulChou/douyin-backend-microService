package service

import (
	"context"
	"crypto/md5"
	"douyin_backend_microService/pkg/errno"
	"douyin_backend_microService/user/dal/db"
	"douyin_backend_microService/user/kitex_gen/userdemo"
	"douyin_backend_microService/user/pack"

	"fmt"
	"io"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func (s *UserService) MGetUser(request *userdemo.MGetUserRequest) ([]*userdemo.User, error) {
	users, err := db.MGetUsers(s.ctx, request.GetUserIds())
	if err != nil {
		return nil, err
	}
	return pack.ConcertUsers(users), nil
}

func (s *UserService) GetUser(request *userdemo.GetUserRequest) (*db.User, error) {
	users, err := db.QueryUserByID(request.UserId)
	if err != nil {
		return nil, err
	}
	return users, nil

}

func (s *UserService) CreateUser(request *userdemo.CreateUserRequest) (int64, error) {
	users, err := db.QueryUser(request.Name)

	if err != nil {
		return 0, errno.ServiceErr
	}

	if len(users) != 0 {
		return 0, errno.UserAlreadyExistErr
	}

	hash := md5.New()
	if _, err = io.WriteString(hash, request.Password); err != nil {
		return 0, err
	}
	password := fmt.Sprintf("%x", hash.Sum(nil))

	return db.CreateUser(s.ctx, &db.User{Password: password, UserName: request.Name})
}

// CheckUser check user info
func (s *UserService) CheckUser(req *userdemo.CheckUserRequest) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.Name
	users, err := db.QueryUser(userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}
	u := users[0]
	if u.Password != passWord {
		return 0, errno.AuthorizationFailedErr
	}
	return int64(u.ID), nil
}

func (s *UserService) UpdateUserFollows(userid int64, toUserId int64, count int64) error {
	return db.UpdateUserFollows(userid, toUserId, count)
}
