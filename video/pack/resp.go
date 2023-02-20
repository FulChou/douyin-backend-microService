package pack

import (
	"douyin_backend_microService/pkg/errno"
	"douyin_backend_microService/video/kitex_gen/videodemo"
	"time"
)

func baseResp(err errno.ErrNo) *videodemo.BaseResp {
	return &videodemo.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg, ServiceTime: time.Now().Unix()}
}

func BuildResponeseMessage(err error) *videodemo.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	errNo := errno.ErrNo{}
	errNo = errno.ConvertErr(err)
	return baseResp(errNo)
}
