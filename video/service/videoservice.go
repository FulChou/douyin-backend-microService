package service

import (
	"context"
	"douyin_backend_microService/user/kitex_gen/userdemo"
	"douyin_backend_microService/video/dal/db"
	"douyin_backend_microService/video/kitex_gen/videodemo"
	"douyin_backend_microService/video/pack"
	"douyin_backend_microService/video/rpc"
	"time"
)

type videoservice struct {
	ctx context.Context
}

func NewVideoService(ctx context.Context) (service *videoservice) {
	return &videoservice{ctx: ctx}
}

func (s *videoservice) PublishList(ctx context.Context, userId uint) ([]*videodemo.Video, error) {
	vs, err := db.VideoListByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}
	videos := pack.ConcertVideos(vs)
	ids := make([]int64, 0)
	uidtovid := make(map[int64]int64, 0)
	for _, v := range vs {
		ids = append(ids, int64(v.UserId))
		uidtovid[int64(v.ID)] = int64(v.UserId)
	}
	users, err := rpc.MGetUser(ctx, &userdemo.MGetUserRequest{UserIds: ids})
	if err != nil {
		return nil, err
	}

	for _, v := range videos {
		uid := uidtovid[int64(v.Id)]
		userinfo := users[uid]
		v.Author = &videodemo.User{
			Id:            userinfo.Id,
			Name:          userinfo.Name,
			FollowCount:   userinfo.FollowCount,
			FollowerCount: userinfo.FollowerCount,
			IsFollow:      userinfo.IsFollow,
		}

	}
	return videos, nil
}

func (s *videoservice) PublishVideo(ctx context.Context, request *videodemo.PublishRequest) error {
	return db.PublishVideo(ctx, &db.Video{
		UserId:        uint(request.UserId),
		Title:         request.Title,
		PlayUrl:       request.Playurl,
		CoverUrl:      request.Coverurl,
		FavoriteCount: 0,
		CommentCount:  0,
	})
}

func (s *videoservice) FeedVideo(ctx context.Context, latestTime time.Time) ([]*videodemo.Video, int64, error) {
	vs, err, nexttime := db.FeedVideos(ctx, latestTime)
	if err != nil {
		return nil, 0, err
	}
	videos := pack.ConcertVideos(vs)
	ids := make([]int64, 0)
	uidtovid := make(map[int64]int64, 0)
	for _, v := range vs {
		ids = append(ids, int64(v.UserId))
		uidtovid[int64(v.ID)] = int64(v.UserId)
	}
	users, err := rpc.MGetUser(ctx, &userdemo.MGetUserRequest{UserIds: ids})
	if err != nil {
		return nil, 0, err
	}

	for _, v := range videos {
		uid := uidtovid[int64(v.Id)]
		userinfo := users[uid]
		v.Author = &videodemo.User{
			Id:            userinfo.Id,
			Name:          userinfo.Name,
			FollowCount:   userinfo.FollowCount,
			FollowerCount: userinfo.FollowerCount,
			IsFollow:      userinfo.IsFollow,
		}

	}

	return videos, nexttime.Unix(), nil

}
