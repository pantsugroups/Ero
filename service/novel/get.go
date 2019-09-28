package novel

import (
	"eroauz/models"
	"eroauz/serializer"
)

type GetService struct {
	ID uint `json:"id" form:"id" null:"false"`
}

func (service *GetService) Get() (models.Novel, *serializer.Response) {
	var novel models.Novel
	if err := models.DB.Where("ID = ?", service.ID).First(&novel).Error; err != nil {
		return novel, &serializer.Response{
			Status: 40003,
			Msg:    "获取失败",
		}
	}
	return novel, nil
}
