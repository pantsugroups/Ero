package volume

import (
	"eroauz/models"
	"eroauz/serializer"
)

type ListService struct {
	ID       uint `json:"id" param:"id"`
	Page     int  `json:"page" form:"page" query:"page"`
	Limit    int  `json:"limit" form:"limit" query:"limit"`
	Offset   int  `json:"offset" form:"offset" query:"offset"`
	PageSize int  `json:"page_count" form:"page_Size" query:"page_Size"`
	Count    int  // 查询结果请求
	All      int  //总数
	result   []models.Volume
}

// 判断是否有上一页或者下一页
func (service *ListService) HaveNextOrLast() (next bool, last bool) {
	if service.Page <= 1 {
		last = false
	} else {
		last = true
	}
	if service.All-(service.Page+1)*service.PageSize >= 0 {
		next = true
	} else if -(service.All - (service.Page+1)*service.PageSize) <= service.PageSize {
		next = true
	} else {
		next = false
	}
	return next, last

}

// 返回查询结果总页数,是按照当前请求的结果的数量除以总数得出的
func (service *ListService) Pages() (int, *serializer.Response) {
	var count int
	if err := models.DB.Model(&models.Volume{}).Count(&service.All).Error; err != nil {
		return 0, &serializer.Response{
			Status: 500,
			Msg:    "查询总数失败",
		}
	}

	if service.Count == 0 {
		return 0, nil
	}
	count = service.All / service.PageSize
	if service.All%service.PageSize != 0 {
		count += 1
	}
	return count, nil
}

// EroAPI godoc
// @Summary 分卷列表
// @Description
// @Tags volume
// @Accept html
// @Produce json
// @Success 200 {object} serializer.NovelListResponse
// @Failure 500 {object} serializer.Response
// @Router /api/v1/novel/volume/:id [get]
// @Param id formData integer true "小说ID"
// @Param page formData integer false "Pages"
// @Param page formData integer false "Pages"
// @Param limit formData integer false "Limit"
// @Param offset formData integer false "Offset"
// @Param page_size formData int false "PageSize default is 10"
func (service *ListService) Pull(create uint) *serializer.Response {
	var volume []models.Volume
	//var count int
	if service.PageSize == 0 {
		service.PageSize = 10
	}

	DB := models.DB.Where("novel_id = ?", service.ID)

	if service.Page > 0 && service.PageSize > 0 {
		DB = DB.Limit(service.PageSize).Offset((service.Page - 1) * service.PageSize)
	} else {
		if service.Limit != 0 {
			DB.Limit(service.Limit)
		}
		if service.Offset != 0 {
			DB.Offset(service.Offset)
		}
	}
	if err := DB.Find(&volume).Order("z_index desc").Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "获取失败",
		}
	}
	service.Count = len(volume)
	for i := range volume {
		if volume[i].NovelID != 0 {
			volume[i].Novel, _ = models.GetNovel(volume[i].Novel)

		}
		if volume[i].FileID != 0 {
			volume[i].File, _ = models.GetFile(volume[i].FileID)
		}
	}
	service.result = volume
	return nil
}
func (service *ListService) Counts() int {
	return service.Count
}
func (service *ListService) Response() interface{} {
	next, last := service.HaveNextOrLast()
	var pages int
	var err *serializer.Response
	if pages, err = service.Pages(); err != nil {
		return err
	}
	return serializer.BuildVolumeListResponse(service.result, service.All, service.Count, next, last, pages)
}
