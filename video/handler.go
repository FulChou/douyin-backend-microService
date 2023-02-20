package main

import (
	"context"
	"douyin_backend_microService/pkg/errno"
	vediodemo "douyin_backend_microService/video/kitex_gen/videodemo"
	"douyin_backend_microService/video/pack"
	"douyin_backend_microService/video/service"
	"time"
)

func UnixSecondToTime(second int64) time.Time {
	return time.Unix(second, 0)
}

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *vediodemo.FeedRequest) (resp *vediodemo.FeedResponse, err error) {
	// TODO: Your code here...
	var Lasttime time.Time
	videoService := service.NewVideoService(ctx)
	if req.LatestTime == 0 {
		Lasttime = time.Now()
	} else {
		Lasttime = UnixSecondToTime(req.LatestTime)
	}

	videos, nexttime, err := videoService.FeedVideo(ctx, Lasttime)
	resp = new(vediodemo.FeedResponse)
	if err != nil {
		resp.BaseResp = pack.BuildResponeseMessage(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildResponeseMessage(errno.Success)
	resp.VideoList = videos

	resp.NextTime = nexttime

	return resp, nil
}

// Publish implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Publish(ctx context.Context, req *vediodemo.PublishRequest) (resp *vediodemo.PublishResponse, err error) {
	// TODO: Your code here...
	if req.UserId < 0 || len(req.Title) == 0 || req.Coverurl == "" || req.Playurl == "" {
		return nil, errno.ParamErr
	}
	err = service.NewVideoService(ctx).PublishVideo(ctx, req)
	if err != nil {
		resp.BaseResp = pack.BuildResponeseMessage(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildResponeseMessage(errno.Success)

	return resp, nil
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *vediodemo.PublishListRequest) (resp *vediodemo.PublishListResponse, err error) {
	// TODO: Your code here...
	resp = new(vediodemo.PublishListResponse)
	if req.UserId <= 0 {
		pack.BuildResponeseMessage(errno.ParamErr)
		resp.BaseResp = pack.BuildResponeseMessage(errno.ParamErr)
		return resp, err
	}

	videoModels, err := service.NewVideoService(ctx).PublishList(ctx, uint(req.UserId))
	if err != nil {
		resp.BaseResp = pack.BuildResponeseMessage(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildResponeseMessage(errno.Success)

	resp.Videos = videoModels
	return resp, nil

}
