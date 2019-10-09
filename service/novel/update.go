package novel

import (
	"eroauz/models"
	"eroauz/serializer"
)

type UpdateService struct {
	ID     uint   `json:"id" form:"id" param:"id" null:"false"`
	Title  string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
	Cover  string `json:"cover" form:"cover"`
	Ended  bool   `json:"ended" form:"ended"`
	Level  int    `json:"level" form:"level"`
	result models.Novel
}

// EroAPI godoc
// @Summary 更新小说
// @Description
// @Tags novel
// @Accept html
// @Produce json
// @Success 200 {object} serializer.NovelResponse
// @Failure 500 {object} serializer.Response
// @Param id path int false "小说ID"
// @Param title formData string true "小说标题"
// @Param cover formData string true "小说封面"
// @Param Ended formData bool true "是否完结"
// @Param level formData int true "目标等级：还未实现"
// @Router /api/v1/novel/:id [put]
// @Security ApiKeyAuth
func (service *UpdateService) Update(create uint) *serializer.Response {
	var novel models.Novel
	if err := models.DB.Where("ID = ?", service.ID).First(&novel).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "获取失败",
		}
	}
	if err := models.DB.Model(&novel).Update(models.Novel{
		Title:  service.Title,
		Author: service.Author,
		Cover:  service.Cover,
		Ended:  service.Ended,
		Level:  service.Level,
	}); err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "获取失败",
		}
	}
	service.result = novel
	return nil
}
func (service *UpdateService) Response() interface{} {
	return serializer.BuildNovelResponse(service.result)
}
