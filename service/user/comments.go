package user

import (
	"eroauz/models"
	"eroauz/serializer"
)

type CommentListService struct {
	Page     int `json:"page" form:"page" query:"page"`
	Limit    int `json:"limit" form:"limit" query:"limit"`
	Offset   int `json:"offset" form:"offset" query:"offset"`
	PageSize int `json:"page_count" form:"page_Size" query:"page_Size"`
	Count    int // 查询结果请求
	All      int //总数
	result   []models.Comment
}

// 判断是否有上一页或者下一页
func (service *CommentListService) HaveNextOrLast() (next bool, last bool) {
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
func (service *CommentListService) Pages() (int, *serializer.Response) {

	if int(service.Count) == 0 {
		return 0, nil
	}
	return int(service.All / service.Count), nil
}

// EroAPI godoc
// @Summary 获取用户发送的评论列表
// @Description
// @Tags comment,user
// @Accept html
// @Produce json
// @Success 200 {object} serializer.CommentListResponse
// @Failure 500 {object} serializer.Response
// @Router /api/v1/user/comments/ [get]
// @Param page formData integer false "Pages"
// @Param limit formData integer false "Limit"
// @Param offset formData integer false "Offset"
// @Param page_size formData integer false "PageSize default is 10"
func (service *CommentListService) Pull(create uint) *serializer.Response {
	var comments []models.Comment
	//var count int
	if service.PageSize == 0 {
		service.PageSize = 10
	}

	DB := models.DB.Where("author_id = ?", create)

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
	if err := DB.Find(&comments).Count(&service.Count).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "获取失败",
		}
	}
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
	if err := models.DB.Model(&models.Comment{}).Where("author_id = ?", create).Count(&service.All).Error; err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "查询总数失败",
		}
	}
	return nil
}
func (service *CommentListService) Counts() int {
	return service.Count
}
func (service *CommentListService) Response() interface{} {
	next, last := service.HaveNextOrLast()
	var pages int
	var err *serializer.Response
	if pages, err = service.Pages(); err != nil {
		return err
	}
	return serializer.BuildCommentListResponse(service.result, service.Count, next, last, pages)
}
