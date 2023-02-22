package service

import (
	"context"
	"douyin_backend_microService/message/dal/db"
	"errors"
)

type MessageService struct {
	ctx context.Context
}

func NewMessageService(ctx context.Context) *MessageService {
	return &MessageService{ctx: ctx}
}

func (s MessageService) MessageAction(userID int64, toUserId int64, actionType int32, messageText string) error {
	if actionType == 1 {
		if err := db.CreateMessage(s.ctx, db.Message{UserID: uint(userID), ToUserID: uint(toUserId), Content: messageText}); err != nil {
			return errors.New("create this comment raise error in db")
		}
	} else {
		return errors.New("not support this action_type")
	}
	return nil
}

func (s MessageService) GetMessageList(userId int64, toUserID int64) ([]*db.Message, error) {
	messageList, err := db.GetMessagesByUserID(uint(userId), uint(toUserID))
	if err != nil {
		return nil, errors.New("fail in finding chat messages")
	}
	return messageList, nil
}
