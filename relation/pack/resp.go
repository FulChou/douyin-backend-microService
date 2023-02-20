package pack

import (
	"douyin_backend_microService/pkg/errno"
	"douyin_backend_microService/relation/kitex_gen/relationdemo"
	"time"
)

func baseResp(err errno.ErrNo) *relationdemo.BaseResp {
	return &relationdemo.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg, ServiceTime: time.Now().Unix()}
}

func BuildResponeseMessage(err error) *relationdemo.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	errNo := errno.ErrNo{}
	errNo = errno.ConvertErr(err)
	return baseResp(errNo)
}
