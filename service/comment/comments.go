package comment

import (
	"eroauz/models"
	"eroauz/serializer"
)

type ListService struct {
	ID       int `json:"id"   form:"id" param:"id" query:"id" null:"false"`
	Type     int `json:"type" query:"type" form:"type"`
	Page     int `json:"page" form:"page" query:"page"`
	Limit    int `json:"limit" form:"limit" query:"limit"`
	Offset   int `json:"offset" form:"offset" query:"offset"`
	PageSize int `json:"page_count" form:"page_Size" query:"page_Size"`
	Count    int // 查询结果请求
	All      int //总数
	result   []models.Comment
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
	if err := models.DB.Model(&models.Comment{}).Count(&service.All).Error; err != nil {
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
// @Summary 评论列表
// @Description
// @Tags comment
// @Accept html
// @Produce json
// @Success 200 {object} serializer.CommentListResponse
// @Failure 500 {object} serializer.Response
// @Router /api/v1/comment/ [get]
// @Param id path integer false "回复对象ID"
// @Param type formData integer false "类型：1为文章，2为小说"
// @Param page formData integer false "Pages"
// @Param limit formData integer false "Limit"
// @Param offset formData integer false "Offset"
// @Param page_size formData integer false "PageSize default is 10"
func (service *ListService) Pull(create uint) *serializer.Response {

	var comments []models.Comment
	//var count int
	if service.PageSize == 0 {
		service.PageSize = 10
	}

	DB := models.DB.Where("type = ?", service.Type).Where("r_id = ?", service.ID)

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
	if err := DB.Find(&comments).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "获取失败",
		}
	}
	service.Count = len(comments)
	for c := range comments {
		var user models.User
		u, err := models.GetUser(comments[c].AuthorID)
		user = u
		if err != nil {
			user.Nickname = "已删除用户"
		}
		comments[c].Author = user

	}
	service.result = comments
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
	return serializer.BuildCommentListResponse(service.result, service.All, service.Count, next, last, pages)
}
