package service

import (
	"context"
	"douyin_backend_microService/relation/dal/db"
	"douyin_backend_microService/relation/rpc"
	"douyin_backend_microService/user/kitex_gen/userdemo"
)

type RelationService struct {
	ctx context.Context
}

func NewRelationService(ctx context.Context) *RelationService {
	return &RelationService{ctx: ctx}
}

func (s *RelationService) CreateFollow(userid int64, touserid int64) error {
	err := db.CreateRelation(s.ctx, db.Follow{
		UserID:       uint(userid),
		FollowUserID: uint(touserid),
	})

	if err != nil {
		return err
	}
	err = rpc.UpdateFollow(s.ctx, &userdemo.UpdateUserFollowRequest{
		UserId:   userid,
		ToUserId: touserid,
		Count:    1,
	})
	if err != nil {
		return err
	}
	return nil

}
func (s *RelationService) ConcealFollow(userid int64, touserid int64) error {
	err := db.DeleteRelation(s.ctx, db.Follow{
		UserID:       uint(userid),
		FollowUserID: uint(touserid),
	})

	if err != nil {
		return err
	}
	err = rpc.UpdateFollow(s.ctx, &userdemo.UpdateUserFollowRequest{
		UserId:   userid,
		ToUserId: touserid,
		Count:    -1,
	})
	if err != nil {
		return err
	}
	return nil

}

func (s *RelationService) GetRelationFollowList(userid int64) ([]*db.Follow, error) {
	return db.GetFollowList(s.ctx, uint(userid))
}

func (s *RelationService) GetRelationFollowerList(userid int64) ([]*db.Follow, error) {
	return db.GetFollowerList(s.ctx, uint(userid))
}

func (s *RelationService) GetFriendsList(userid int64) ([]int64, error) {
	list, err := db.GetFriendList(s.ctx, uint(userid))
	if err != nil {
		return nil, err
	}
	friendIds := make([]int64, 0)
	for _, frend := range list {
		friendIds = append(friendIds, int64(frend.UserID))
	}
	return friendIds, nil
}

// 获取关注博主id
func (s *RelationService) GetFollowUserIds(follows []*db.Follow) []int64 {
	followuserIds := make([]int64, 0)
	for _, follow := range follows {
		followuserIds = append(followuserIds, int64(follow.FollowUserID))
	}
	return followuserIds

}

// 获取粉丝id
func (*RelationService) GetFollowerUserIds(follows []*db.Follow) []int64 {
	followuserIds := make([]int64, 0)
	for _, follow := range follows {
		followuserIds = append(followuserIds, int64(follow.UserID))
	}
	return followuserIds

}
