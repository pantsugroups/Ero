package comment

import (
	"eroauz/models"
	"eroauz/serializer"
)

type GetService struct {
	ID     uint `json:"id" form:"id" param:"id" null:"false"`
	result models.Comment
}

func (service *GetService) Get(create uint) *serializer.Response {
	var comment models.Comment
	if err := models.DB.Where("ID = ?", service.ID).First(&comment).Error; err != nil {
		return &serializer.Response{
			Status: 40003,
			Msg:    "获取失败",
			Error:  err.Error(),
		}
	}
	var user models.User
	u, err := models.GetUser(comment.AuthorID)
	user = u
	if err != nil {
		user.Nickname = "已删除用户"
	}
	comment.Author = user
	service.result = comment
	return nil
}
func (service *GetService) Response() interface{} {
	return serializer.BuildCommentResponse(service.result)

}
