package serializer

import (
	"eroauz/models"
	"time"
)

type Message struct {
	ID       uint      `json:"id"`
	Title    string    `json:"title"`
	SendName string    `json:"send_name"`
	SendID   uint      `json:"send_id"`
	RecvName string    `json:"recv_name"`
	RecvID   uint      `json:"recv_id"`
	CreateAt time.Time `json:"create_time"`
	Read     bool      `json:"is_read"`
}

type MessageResponse struct {
	Response
	Data Message `json:"data"`
}

type MessageListResponse struct {
	Response
	Count int       `json:"count"`
	All   int       `json:"all"`
	Data  []Message `json:"data"`
	Next  bool      `json:"have_next"`
	Last  bool      `json:"have_last"`
	Pages int       `json:"pages"`
}

func BuildMessage(message models.Message) Message {
	return Message{
		ID:       message.ID,
		Title:    message.Title,
		SendID:   message.Send.ID,
		SendName: message.Send.Nickname,
		RecvID:   message.Recv.ID,
		RecvName: message.Recv.Nickname,
		CreateAt: message.CreatedAt,
		Read:     message.Read,
	}
}

func BuildMessageList(messages []models.Message) []Message {
	var messageList []Message
	for _, a := range messages {
		i := BuildMessage(a)
		messageList = append(messageList, i)
	}
	return messageList
}

func BuildMessageResponse(message models.Message) MessageResponse {
	return MessageResponse{
		Data: BuildMessage(message),
	}
}

func BuildMessageListResponse(messages []models.Message, all int, count int, next bool, last bool, pages int) MessageListResponse {
	return MessageListResponse{
		Count: count,
		All:   all,
		Data:  BuildMessageList(messages),
		Next:  next,
		Last:  last,
		Pages: pages,
	}
}
