package relationship

import (
	"eroauz/models"
	"eroauz/serializer"
)

type ArchiveListService struct {
	ID       int `json:"id" form:"id" query:"id" param:"id"`
	Page     int `json:"page" form:"page" query:"page"`
	Limit    int `json:"limit" form:"limit" query:"limit"`
	Offset   int `json:"offset" form:"offset" query:"offset"`
	PageSize int `json:"page_count" form:"page_Size" query:"page_Size"`
	Count    int // 查询结果请求
	All      int //总数
	result   []models.Archive
}

// 判断是否有上一页或者下一页
func (service *ArchiveListService) HaveNextOrLast() (next bool, last bool) {
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

func (service *ArchiveListService) Pages() (int, *serializer.Response) {
	var count int
	if err := models.DB.Model(&models.ArchiveCategory{}).Count(&service.All).Error; err != nil {
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
// @Summary 获取分类下的小说列表
// @Description
// @Tags archive
// @Accept html
// @Produce json
// @Success 200 {object} serializer.ArchiveListResponse
// @Failure 500 {object} serializer.Response
// @Router /api/v1/category/archives/:id [get]
// @Param id path integer false "id"
// @Param page formData integer false "Pages"
// @Param limit formData integer false "Limit"
// @Param offset formData integer false "Offset"
// @Param page_size formData integer false "PageSize default is 10"
func (service *ArchiveListService) Pull(create uint) *serializer.Response {
	var categories []models.ArchiveCategory
	//var count int

	if service.PageSize == 0 {
		service.PageSize = 10
	}

	DB := models.DB

	if service.Page > 0 && service.PageSize > 0 {
		DB = DB.Limit(service.PageSize).Offset((service.Page-1)*service.PageSize).Where("pass = ?", true).Order("updated_at")
	} else {
		if service.Limit != 0 {
			DB.Limit(service.Limit)
		}
		if service.Offset != 0 {
			DB.Offset(service.Offset)
		}
	}
	if err := DB.Find(&categories).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "获取失败",
		}
	}
	service.Count = len(categories)
	for a := range categories {
		var archive models.Archive
		var user models.User
		archive, _ = models.GetArchive(categories[a].ArchiveID)
		u, err := models.GetUser(archive.Create)
		user = u
		if err != nil {
			user.Nickname = "被删除用户"
		}
		service.result = append(service.result, archive)
	}

	return nil
}
func (service *ArchiveListService) Counts() int {
	return service.Count
}
func (service *ArchiveListService) Response() interface{} {
	next, last := service.HaveNextOrLast()
	var pages int
	var err *serializer.Response
	if pages, err = service.Pages(); err != nil {
		return err
	}
	return serializer.BuildArchiveListResponse(service.result, service.All, service.Count, next, last, pages)
}
