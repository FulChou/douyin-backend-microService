package db

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName         string `json:"user_name"`
	Password         string `json:"password"`
	Avatar           string `json:"avatar_url"`
	background_image string `json:"backgroundImage_url"`
	signature        string `json:"signature"`
	FollowCount      uint64 `json:"follow_count" `
	FollowerCount    uint64 `json:"follower_count"`
}

func (u *User) TableName() string {
	return UserTableName
}

// CreateUser create user info
func CreateUsers(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}
func CreateUser(ctx context.Context, user *User) (int64, error) {
	return int64(user.ID), DB.WithContext(ctx).Create(user).Error
}

// QueryUser query list of user info by name
func QueryUser(userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryUserByID(ID int64) (*User, error) {
	var res *User
	if err := DB.Where("id = ?", ID).First(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func CheckUser(account, password string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.Where("user_name = ?", account).Where("password = ?", password).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateUserFollows(userId int64, toUserId int64, count int64) error {
	// update My_follow_count
	user, err := QueryUserByID(userId)
	if err != nil {
		return errors.New("user doesn't exist in db")
	}
	if count == -1 && user.FollowCount == 0 {
		return errors.New("follow_count already zero")
	}

	if err := DB.Model(&User{}).Where("id = ?", userId).Update("follow_count", int(user.FollowCount)+int(count)).Error; err != nil {
		return errors.New("update user follow_count failed")
	}

	// update follower_count
	toUser, err := QueryUserByID(toUserId)
	if err != nil {
		return errors.New("toUser doesn't exist in db")
	}
	if count == -1 && user.FollowerCount == 0 {
		return errors.New("follower_count already zero")
	}
	if err := DB.Model(&User{}).Where("id = ?", toUserId).Update("follower_count", int(toUser.FollowerCount)+int(count)).Error; err != nil {
		return errors.New("update user follower_count failed")
	}
	return nil
}
