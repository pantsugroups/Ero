package archive

import (
	"eroauz/models"
	"eroauz/serializer"
)

type GetService struct {
	ID     uint `json:"id" form:"id" param:"id" null:"false"`
	result models.Archive
}

// EroAPI godoc
// @Summary 获取文章
// @Description 获取单个文章的详细信息
// @Tags archive
// @Accept html
// @Produce json
// @Success 200 {object} serializer.ArchiveResponse
// @Failure 500 {object} serializer.Response
// @Param id path integer true "文章ID"
// @Router /api/v1/archive/:id [get]
func (service *GetService) Get(create uint) *serializer.Response {
	var archive models.Archive
	if err := models.DB.Where("ID = ?", service.ID).First(&archive).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "获取失败",
			Error:  err.Error(),
		}
	}
	var user models.User
	u, err := models.GetUser(archive.Create)
	user = u
	if err != nil {
		user.Nickname = "被删除用户"
	}
	archive.Create = user
	service.result = archive
	return nil
}
func (service *GetService) Response() interface{} {
	return serializer.BuildArchiveResponse(service.result)

}
