package archive

import (
	"eroauz/models"
	"eroauz/serializer"
)

type GetService struct{
	ID uint `json:"id" form:"id" null:"false"`
}
func (service *GetService)Get()(models.Archive,*serializer.Response){
	var archive models.Archive
	if err:=models.DB.Where("ID = ?",service.ID).First(&archive).Error;err!=nil{
		return archive, &serializer.Response{
			Status: 40003,
			Msg:    "获取失败",
		}
	}
	return archive,nil
}