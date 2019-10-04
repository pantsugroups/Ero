package category

import (
	"eroauz/models"
	"eroauz/serializer"
)

type UpdateService struct {
	ID     uint   `json:"id" form:"id" param:"id" null:"false"`
	Title  string `json:"title" form:"title"`
	result models.Category
}

// EroAPI godoc
// @Summary 创建分类
// @Description 必须为管理员
// @Tags category,admin
// @Accept html
// @Produce json
// @Success 200 {object} serializer.CategoryResponse
// @Failure 500 {object} serializer.Response
// @Param id path int false "分类ID"
// @Param title formData string true "分类标题"
// @Router /api/v1/category/:id [put]
// @Security ApiKeyAuth
func (service *UpdateService) Update(create uint) *serializer.Response {
	var category models.Category
	if err := models.DB.Where("ID = ?", service.ID).First(&category).Error; err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
	if err := models.DB.Model(&category).Update(models.Novel{
		Title: service.Title,
	}); err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
	service.result = category
	return nil
}
func (service *UpdateService) Response() interface{} {
	return serializer.BuildCategoryResponse(service.result)
}
