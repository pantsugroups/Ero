package service
// 试图骚操作，但是失败了.jpg
//
//import (
//	"eroauz/models"
//	"eroauz/serializer"
//)
//
//
//
//type ListService struct{
//	Page int `json:"page" form:"page" query:"page"`
//	Limit  int `json:"limit" form:"limit" query:"limit"`
//	Offset int `json:"offset" form:"offset" query:"offset"`
//	PageSize int `json:"page_count" form:"page_Size" query:"page_Size"`
//	Count int // 查询结果请求
//	All   int//总数
//}
//// 判断是否有上一页或者下一页
//func (service *ListService)HaveNextOrLast()(next bool,last bool){
//	if service.Page<=1{
//		last = false
//	}else{
//		last = true
//	}
//	if service.All - (service.Page+1)*service.PageSize<0{
//		next = false
//	}else{
//		next = true
//	}
//	return next,last
//
//}
//
//// 返回查询结果总页数,是按照当前请求的结果的数量除以总数得出的
//func (service *ListService)Pages()(int,*serializer.Response){
//
//	if err:=models.DB.Model(&models.Archive{}).Count(&service.All).Error;err != nil{
//		return 0,&serializer.Response{
//			Status: 40005,
//			Msg:    "查询总数失败",
//		}
//	}
//	return int(service.All/service.Count),nil
//}
//func (service *ListService)Pull(modes *[]interface{})*serializer.Response{
//
//	//var count int
//	if service.PageSize == 0 {
//		service.PageSize = 10
//	}
//
//	DB :=models.DB
//
//	if service.Page > 0 && service.PageSize > 0 {
//		DB = DB.Limit(service.Page).Offset((service.Page - 1) * service.PageSize)
//	}else{
//		if service.Limit != 0{
//			DB.Limit(service.Limit)
//		}
//		if service.Offset != 0{
//			DB.Offset(service.Offset)
//		}
//	}
//	if err := DB.Find(modes).Count(&service.Count).Error;err != nil{
//		return  &serializer.Response{
//			Status: 40005,
//			Msg:    "获取失败",
//		}
//	}
//	return nil
//}