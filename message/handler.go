package main

import (
	"context"
	messagedemo "douyin_backend_microService/message/kitex_gen/messagedemo"
	"douyin_backend_microService/message/pack"
	"douyin_backend_microService/message/service"
	"douyin_backend_microService/pkg/errno"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// GetMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetMessage(ctx context.Context, req *messagedemo.MessageRequest) (resp *messagedemo.MessageResponse, err error) {
	// TODO: Your code here...
	resp = new(messagedemo.MessageResponse)
	messageService := service.NewMessageService(ctx)
	messageList, err := messageService.GetMessageList(req.UserId, req.ToUserId)

	if err != nil {
		resp.BaseResp = pack.BuildResponeseMessage(err)
		return resp, err
	}
	res := make([]*messagedemo.Message, 0)
	for _, v := range messageList {
		res = append(res, &messagedemo.Message{Id: int64(v.ID), ToUserId: int64(v.ToUserID),
			FromUserId: int64(v.UserID), Content: v.Content, CreateTime: v.CreatedAt.Format("2006-01-02 15:04:05")})
	}
	resp.MessageList = res

	resp.BaseResp = pack.BuildResponeseMessage(errno.Success)
	return resp, nil

}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *messagedemo.MessageActionRequest) (resp *messagedemo.MessageActionResponse, err error) {
	// TODO: Your code here...
	resp = new(messagedemo.MessageActionResponse)
	messageService := service.NewMessageService(ctx)
	err = messageService.MessageAction(req.UserId, req.ToUserId, req.ActionType, req.Content)
	if err != nil {
		resp.BaseResp = pack.BuildResponeseMessage(err)
	}
	resp.BaseResp = pack.BuildResponeseMessage(errno.Success)
	return resp, nil
}
