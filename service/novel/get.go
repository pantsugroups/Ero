package novel

import (
	"eroauz/models"
	"eroauz/serializer"
)

type GetService struct {
	ID     uint `json:"id" form:"id" param:"id" null:"false"`
	result models.Novel
}

// EroAPI godoc
// @Summary 获取小说详细
// @Description
// @Tags novel
// @Accept html
// @Produce json
// @Success 200 {object} serializer.NovelResponse
// @Failure 500 {object} serializer.Response
// @Param id path integer true "小说ID"
// @Router /api/v1/novel/:id [get]
func (service *GetService) Get(create uint) *serializer.Response {
	var novel models.Novel
	if err := models.DB.Where("ID = ?", service.ID).First(&novel).Error; err != nil {
		return &serializer.Response{
			Status: 40003,
			Msg:    "获取失败",
		}
	}
	var user models.User
	u, err := models.GetUser(novel.Create)
	user = u
	if err != nil {
		user.Nickname = "已删除用户"
	}
	novel.Create = user
	service.result = novel
	return nil
}
func (service *GetService) Response() interface{} {
	return serializer.BuildNovelResponse(service.result)
}
