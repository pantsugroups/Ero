package volume

import (
	"eroauz/models"
	"eroauz/serializer"
)

type CreateService struct {
	ID     uint   `json:"id" param:"id" form:"id" query:"id"`
	Title  string `json:"title" form:"title" query:"title"`
	Cover  string `json:"cover" form:"cover" query:"cover"`
	File   int    `json:"file" form:"file"`
	result models.Volume
}

func (service *CreateService) Create(create uint) *serializer.Response {
	var file models.File

	n, err := models.GetNovel(service.ID)
	if err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "寻找匹配ID失败",
		}
	}
	if service.File != 0 {
		f, err := models.GetFile(service.File)
		if err != nil {
			return &serializer.Response{
				Status: 40005,
				Msg:    "寻找匹配ID失败",
			}
		}
		file = f
	}
	volume := models.Volume{
		Title: service.Title,
		Cover: service.Cover,
		Novel: n,
		File:  file,
	}
	if err := models.DB.Create(&volume).Error; err != nil {
		return &serializer.Response{
			Status: 40004,
			Msg:    "创建失败",
		}
	}
	service.result = volume
	return nil
}
func (service *CreateService) Response() interface{} {
	return serializer.BuildVolumeResponse(service.result)

}
