package api

// 沙雕写法，已经废弃
//
//import (
//	"eroauz/serializer"
//	"eroauz/service/archive"
//	"fmt"
//	"github.com/labstack/echo"
//)
//func Archive(c echo.Context)(err error){
//	var service archive.ListService
//	var pages int
//	if err := c.Bind(&service); err == nil {
//		if lists, err := service.Pull(); err != nil {
//			return c.JSON(200, err)
//		} else {
//			next,last := service.HaveNextOrLast()
//			if pages,err = service.Pages();err != nil{
//				return c.JSON(200,err)
//			}
//			res := serializer.BuildArchiveListResponse(lists,service.Count,next,last,pages)
//			return c.JSON(200, res)
//		}
//	} else {
//		return c.JSON(200, &serializer.Response{
//			Status: 40001,
//			Msg:    "参数错误",
//			Error:  fmt.Sprint(err),
//		})
//	}
//}
//
//// 获取文章
//func ArchiveGet(c echo.Context) (err error){
//	var service archive.CreateService
//	if err := c.Bind(&service); err == nil {
//		if a, err := service.Create(); err != nil {
//			return c.JSON(200, err)
//		} else {
//			res := serializer.BuildArchiveResponse(a)
//			return c.JSON(200, res)
//		}
//	} else {
//		return c.JSON(200, &serializer.Response{
//			Status: 40001,
//			Msg:    "参数错误",
//			Error:  fmt.Sprint(err),
//		})
//	}
//}
////删除文章
//func ArchiveDelete(c echo.Context) (err error){
//	var service archive.GetService
//	if err := c.Bind(&service); err == nil {
//		if a, err := service.Get(); err != nil {
//			return c.JSON(200, err)
//		} else {
//			res := serializer.BuildArchiveResponse(a)
//			return c.JSON(200, res)
//		}
//	} else {
//		return c.JSON(200, &serializer.Response{
//			Status: 40001,
//			Msg:    "参数错误",
//			Error:  fmt.Sprint(err),
//		})
//	}
//}
////更新文章
//func ArchiveUpdate(c echo.Context) (err error){
//	var service archive.UpdateService
//	if err := c.Bind(&service); err == nil {
//		if a, err := service.Update(); err != nil {
//			return c.JSON(200, err)
//		} else {
//			res := serializer.BuildArchiveResponse(a)
//			return c.JSON(200, res)
//		}
//	} else {
//		return c.JSON(200, &serializer.Response{
//			Status: 40001,
//			Msg:    "参数错误",
//			Error:  fmt.Sprint(err),
//		})
//	}
//}
////新建文章
//func ArchiveNew(c echo.Context)(err error){
//	var service archive.CreateService
//	if err := c.Bind(&service); err == nil {
//		if a, err := service.Create(); err != nil {
//			return c.JSON(200, err)
//		} else {
//			res := serializer.BuildArchiveResponse(a)
//			return c.JSON(200, res)
//		}
//	} else {
//		return c.JSON(200, &serializer.Response{
//			Status: 40001,
//			Msg:    "参数错误",
//			Error:  fmt.Sprint(err),
//		})
//	}
//}
