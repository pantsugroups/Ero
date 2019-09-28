package api

//
//import (
//	"eroauz/serializer"
//	"eroauz/service/archive"
//	"fmt"
//
//	"eroauz/service/novel"
//	"github.com/labstack/echo"
//)
////小说列表
//func Novel(c echo.Context)(err error){
//	var service novel.ListService
//	var pages int
//	if err := c.Bind(&service); err == nil {
//		if lists, err := service.Pull(); err != nil {
//			return c.JSON(200, err)
//		} else {
//			next,last := service.HaveNextOrLast()
//			if pages,err = service.Pages();err != nil{
//				return c.JSON(200,err)
//			}
//			res := serializer.BuildNovelListResponse(lists,service.Count,next,last,pages)
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
//// 获取小说
//func NovelGet(c echo.Context) (err error){
//	var service novel.CreateService
//	if err := c.Bind(&service); err == nil {
//		if a, err := service.Create(); err != nil {
//			return c.JSON(200, err)
//		} else {
//			res := serializer.BuildNovelResponse(a)
//			return c.JSON(200, res)
//		}
//	} else {
//		return c.JSON(200, &serializer.Response{
//			Status: 40001,
//			Msg:    "参数错误",
//			Error:  fmt.Sprint(err),
//		})
//	}
//
//}
////删除小说
//func NovelDelete(c echo.Context) (err error){
//
//
//}
////更新小说
//func NovelUpdate(c echo.Context) (err error){
//
//}
////新建小说
//func NovelNew(c echo.Context)(err error){
//
//
//}
