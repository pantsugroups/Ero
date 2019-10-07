package category

import (
	"eroauz/models"
	"eroauz/serializer"
)

type CreateService struct {
	Title  string `json:"title" form:"title" null:"false"`
	result models.Category
}

// EroAPI godoc
// @Summary 创建分类
// @Description 必须登陆
// @Tags category,admin
// @Accept html
// @Produce json
// @Success 200 {object} serializer.CategoryResponse
// @Failure 500 {object} serializer.Response
// @Param title formData string true "分类标题"
// @Router /api/v1/category/ [post]
// @Security ApiKeyAuth
func (service *CreateService) Create(create uint) *serializer.Response {
	category := models.Category{
		Title: service.Title,
		Count: 0,
	}
	if err := models.DB.Create(&category).Error; err != nil {
		return &serializer.Response{
			Status: 40007,
			Msg:    "创建失败",
		}
	}
	service.result = category
	return nil
}
func (service *CreateService) Response() interface{} {
	return serializer.BuildCategoryResponse(service.result)
}
