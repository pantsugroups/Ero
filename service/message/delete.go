package message

import (
	"eroauz/models"
	"eroauz/serializer"
)

type DeleteService struct {
	ID uint `json:"id" form:"id" param:"id" null:"false"`
}

func (service *DeleteService) Delete(create uint) *serializer.Response {
	var msg models.Message
	if err := models.DB.Where("ID = ?", service.ID).First(&msg).Error; err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "寻找匹配ID失败",
			Error:  err.Error(),
		}
	}
	if msg.RecvID != create {
		u, err := models.GetUser(create)
		if err != nil {
			return &serializer.Response{
				Status: 500,
				Msg:    "内部错误",
				Error:  err.Error(),
			}
		}
		if u.Status != models.Admin {
			return &serializer.Response{
				Status: 403,
				Msg:    "权限不足",
			}
		}
	}
	if err := models.DB.Delete(&msg).Error; err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "删除失败",
			Error:  err.Error(),
		}
	}
	return nil
}
func (service *DeleteService) Response() interface{} {
	return serializer.Response{
		Status: 0,
		Msg:    "成功",
	}
}
