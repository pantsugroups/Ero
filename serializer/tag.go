package serializer

import "eroauz/models"

type Tag struct {
	Title string `json:"title"`
	ID    int    `json:"id"`
	Type  string `json:"Type"`
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
		ID:    tag.RID,
		Type:  models.Int2String_Tag(tag.R_Type),
	}
}

// BuildUserResponse 序列化用户响应
func BuildTagResponse(tag models.Tag) TagResponse {
	return TagResponse{
		Data: BuildTag(tag),
	}
}
