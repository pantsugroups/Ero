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
	if service.All-(service.Page+1)*service.PageSize < 0 {
		next = false
	} else {
		next = true
	}
	return next, last

}

// 返回查询结果总页数,是按照当前请求的结果的数量除以总数得出的
func (service *ListService) Pages() (int, *serializer.Response) {

	if err := models.DB.Model(&models.Volume{}).Count(&service.All).Error; err != nil {
		return 0, &serializer.Response{
			Status: 40005,
			Msg:    "查询总数失败",
		}
	}

	if int(service.Count) == 0 {
		return 0, nil
	}
	return int(service.All / service.Count), nil
}
func (service *ListService) Pull(create uint) *serializer.Response {
	var volume []models.Volume
	//var count int
	if service.PageSize == 0 {
		service.PageSize = 10
	}

	DB := models.DB.Where("novel_id = ?", service.ID)

	if service.Page > 0 && service.PageSize > 0 {
		DB = DB.Limit(service.Page).Offset((service.Page - 1) * service.PageSize)
	} else {
		if service.Limit != 0 {
			DB.Limit(service.Limit)
		}
		if service.Offset != 0 {
			DB.Offset(service.Offset)
		}
	}
	if err := DB.Find(&volume).Count(&service.Count).Error; err != nil {
		return &serializer.Response{
			Status: 40005,
			Msg:    "获取失败",
		}
	}
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
	return serializer.BuildVolumeListResponse(service.result, service.Count, next, last, pages)
}
