package comment

import (
	"eroauz/models"
	"eroauz/serializer"
)

type GetService struct {
	ID     uint `json:"id" form:"id" param:"id" null:"false"`
	result models.Comment
}

// EroAPI godoc
// @Summary 获取评论详细
// @Description 获取单个评论的详细信息
// @Tags comment
// @Accept html
// @Produce json
// @Success 200 {object} serializer.CommentResponse
// @Failure 500 {object} serializer.Response
// @Param id path integer true "评论ID"
// @Router /api/v1/comment/:id [get]
func (service *GetService) Get(create uint) *serializer.Response {
	var comment models.Comment
	if err := models.DB.Where("ID = ?", service.ID).First(&comment).Error; err != nil {
		return &serializer.Response{
			Status: 500,
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
