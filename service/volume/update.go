package volume

import (
	"eroauz/models"
	"eroauz/serializer"
)

type UpdateService struct {
	ID     uint   `json:"id" form:"id" param:"id" null:"false"`
	Title  string `json:"title" form:"title"`
	Cover  string `json:"cover" form:"cover"`
	File   uint   `json:"file" form:"file"`
	Novel  uint   `json:"novel" form:"novel"`
	result models.Volume
}

// EroAPI godoc
// @Summary 更新小说分卷
// @Description 必须登陆
// @Tags volume,admin
// @Accept html
// @Produce json
// @Success 200 {object} serializer.VolumeResponse
// @Failure 500 {object} serializer.Response
// @Param id path int true "分卷ID"
// @Param title formData string true "分卷标题"
// @Param cover formData string false "分卷封面，URL，如果封面为空的话泽会自动替换。默认封面请检查conf.DefaultCover字段"
// @Param file formData integer false "文件ID"
// @Param novel formData integer false "小说ID"
// @Router /api/v1/volume/ [put]
// @Security ApiKeyAuth
func (service *UpdateService) Update(create uint) *serializer.Response {
	var volume models.Volume
	var file models.File
	var novel models.Novel
	if err := models.DB.Where("ID = ?", service.ID).First(&volume).Error; err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
	if service.File != 0 {
		f, err := models.GetFile(service.File)
		if err != nil {
			return &serializer.Response{
				Status: 40005,
				Msg:    "找不到ID",
			}
		}
		file = f
	}
	if service.Novel != 0 {
		n, err := models.GetNovel(service.Novel)
		if err != nil {
			return &serializer.Response{
				Status: 40005,
				Msg:    "找不到ID",
			}
		}
		novel = n
	}
	if err := models.DB.Model(&volume).Update(models.Volume{
		Title: service.Title,
		Cover: service.Title,
		File:  file,
		Novel: novel,
	}); err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
	service.result = volume
	return nil
}
func (service *UpdateService) Response() interface{} {
	return serializer.BuildVolumeResponse(service.result)
}
