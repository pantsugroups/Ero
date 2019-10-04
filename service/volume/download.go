package volume

import (
	"eroauz/models"
	"eroauz/serializer"
	"eroauz/utils"
)

type GetService struct {
	ID     uint `json:"id" form:"id" param:"id" null:"false"`
	result serializer.File
}

// EroAPI godoc
// @Summary 获取分卷下载连接
// @Description 必须登陆,请到/api/v1/download下载
// @Tags volume
// @Accept html
// @Produce json
// @Success 200 {object} serializer.FileResponse
// @Failure 500 {object} serializer.Response
// @Param id path int true "分卷ID"
// @Router /api/v1/volume/:id [get]
// @Security ApiKeyAuth
func (service *GetService) Get(create uint) *serializer.Response {
	var f models.File
	u, err := models.GetUser(create)
	if err != nil {
		return &serializer.Response{
			Status: 403,
			Msg:    "找不到用户",
			Error:  err.Error(),
		}
	}
	if u.Point <= 0 {
		return &serializer.Response{
			Status: 403,
			Msg:    "下载点数不足",
		}
	} else {
		if err := models.DB.Model(&u).Update(models.User{
			Point: u.Point - 1,
		}); err != nil {
			return &serializer.Response{
				Status: 403,
				Msg:    "扣除下载点数失败",
			}
		}
	}
	if err := models.DB.Where("ID = ?", service.ID).First(&f).Error; err != nil {
		return &serializer.Response{
			Status: 40003,
			Msg:    "获取失败",
			Error:  err.Error(),
		}
	}
	file := serializer.BuildFile(f)
	file.Hash = utils.Generate(file.FileName)
	file.Token = utils.RandStringRunes(16)
	service.result = file
	return nil
}
func (service *GetService) Response() interface{} {
	return serializer.BuildDownloadFileResponse(service.result)

}
