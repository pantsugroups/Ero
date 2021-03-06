package comment

import (
	"eroauz/models"
	"eroauz/serializer"
)

type DeleteService struct {
	ID uint `json:"id" form:"id" param:"id" null:"false"`
}

// EroAPI godoc
// @Summary 删除分类
// @Description 必须为管理员
// @Tags comment
// @Accept html
// @Produce json
// @Success 200 {object} serializer.Response
// @Failure 500 {object} serializer.Response
// @Param id path integer true "分类ID"
// @Router /api/v1/comment/:id [delete]
// @Security ApiKeyAuth
func (service *DeleteService) Delete(create uint) *serializer.Response {
	var comment models.Comment
	if err := models.DB.Where("ID = ?", service.ID).First(&comment).Error; err != nil {
		return &serializer.Response{
			Status: 404,
			Msg:    "寻找匹配ID失败",
			Error:  err.Error(),
		}
	}
	if comment.AuthorID != create {
		u, err := models.GetUser(create)
		if err != nil {
			return &serializer.Response{
				Status: 500,
				Msg:    "内部错误",
				Error:  err.Error(),
			}
		}
		if u.Status != models.Admin {
			return &serializer.Response{
				Status: 403,
				Msg:    "权限不足",
			}
		}
	}
	if err := models.DB.Delete(&comment).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "删除失败",
			Error:  err.Error(),
		}
	}
	return nil
}
func (service *DeleteService) Response() interface{} {
	return serializer.Response{
		Status: 0,
		Msg:    "成功",
	}
}
