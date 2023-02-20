package pack

import (
	"douyin_backend_microService/video/dal/db"
	"douyin_backend_microService/video/kitex_gen/videodemo"
)

func ConcertVedio(video *db.Video) (res *videodemo.Video) {
	if video == nil {
		return nil
	}
	return &videodemo.Video{
		Id:       int64(video.ID),
		Author:   nil,
		PlayUrl:  video.PlayUrl,
		CoverUrl: video.CoverUrl,
		Title:    video.Title,
	}
}

func ConcertVideos(vedios []*db.Video) (res []*videodemo.Video) {
	if vedios == nil || len(vedios) == 0 {
		return nil
	}

	res = make([]*videodemo.Video, 0)
	for _, video := range vedios {
		if u := ConcertVedio(video); u != nil {
			res = append(res, u)
		}
	}

	return res
}
