package serializer

import (
	"eroauz/models"
)

type Tag struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

// UserResponse 单个用户序列化
type TagResponse struct {
	Response
	Data Tag `json:"data"`
}

// BuildUser 序列化用户
func BuildTag(tag models.Tag) Tag {
	return Tag{
		Title: tag.Title,
	}
}

// BuildUserResponse 序列化用户响应
func BuildTagResponse(tag models.Tag) TagResponse {
	return TagResponse{
		Data: BuildTag(tag),
	}
}
