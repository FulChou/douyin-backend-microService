package pack

import (
	"douyin_backend_microService/message/kitex_gen/userdemo"
	"douyin_backend_microService/pkg/errno"
	"time"
)

func baseResp(err errno.ErrNo) *userdemo.BaseResp {
	return &userdemo.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg, ServiceTime: time.Now().Unix()}
}

func BuildResponeseMessage(err error) *userdemo.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}
	errNo := errno.ErrNo{}
	errNo = errno.ConvertErr(err)
	return baseResp(errNo)
}
