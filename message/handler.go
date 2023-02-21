package main

import (
	"context"
	message "douyin_backend_microService/message/kitex_gen/message"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// SendMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) SendMessage(ctx context.Context, req *message.MessageActionRequest) (resp *message.MessageActionRequestResponse, err error) {
	// TODO: Your code here...

	return
}

// GetMessageChatRecord implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetMessageChatRecord(ctx context.Context, req *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	// TODO: Your code here...
	return
}
