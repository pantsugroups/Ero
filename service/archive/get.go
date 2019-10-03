package archive

import (
	"eroauz/models"
	"eroauz/serializer"
)

type GetService struct {
	ID     uint `json:"id" form:"id" param:"id" null:"false"`
	result models.Archive
}

func (service *GetService) Get(create uint) *serializer.Response {
	var archive models.Archive
	if err := models.DB.Where("ID = ?", service.ID).First(&archive).Error; err != nil {
		return &serializer.Response{
			Status: 40003,
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
