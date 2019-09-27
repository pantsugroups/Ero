package novel

import (
	"eroauz/models"
	"eroauz/serializer"
)

type DeleteService struct {
	ID uint `json:"ID" form:"ID" null:"false"`
}
func (service *DeleteService)Delete()*serializer.Response{
	var novel models.Novel
	if err:=models.DB.Where("ID = ?",service.ID).First(&novel);err != nil{
		return &serializer.Response{
			Status: 40005,
			Msg:    "寻找匹配ID失败",
		}
	}
	if err := models.DB.Delete(&novel).Error;err!= nil{
		return &serializer.Response{
			Status: 40005,
			Msg:    "删除失败",
		}
	}
	return nil
}