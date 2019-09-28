package api

import (
	"eroauz/interface"
	"eroauz/serializer"
	"fmt"
	"github.com/labstack/echo"
)

// 获取列表
func List(service _interface.ListInterface) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		var pages int
		if err := c.Bind(&service); err == nil {
			if lists, err := service.Pull(); err != nil {
				return c.JSON(200, err)
			} else {
				next, last := service.HaveNextOrLast()
				if pages, err = service.Pages(); err != nil {
					return c.JSON(200, err)
				}
				res := serializer.BuildArchiveListResponse(lists, service.Counts(), next, last, pages)
				return c.JSON(200, res)
			}
		} else {
			return c.JSON(200, &serializer.Response{
				Status: 40001,
				Msg:    "参数错误",
				Error:  fmt.Sprint(err),
			})
		}
	}
}

// 创建
func Create(service _interface.CreateInterface) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		if err := c.Bind(&service); err == nil {
			if err := service.Create(); err != nil {
				return c.JSON(200, err)
			} else {
				res := service.Response()
				return c.JSON(200, res)
			}
		} else {
			return c.JSON(200, &serializer.Response{
				Status: 40001,
				Msg:    "参数错误",
				Error:  fmt.Sprint(err),
			})
		}

	}
}

// 获取单个
func Get(service _interface.GetInterface) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		if err := c.Bind(&service); err == nil {
			if err := service.Get(); err != nil {
				return c.JSON(200, err)
			} else {
				res := service.Response()
				return c.JSON(200, res)
			}
		} else {
			return c.JSON(200, &serializer.Response{
				Status: 40001,
				Msg:    "参数错误",
				Error:  fmt.Sprint(err),
			})
		}
	}
}

//更新
func Update(service _interface.UpdateInterface) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		if err := c.Bind(&service); err == nil {
			if err := service.Update(); err != nil {
				return c.JSON(200, err)
			} else {
				res := service.Response()
				return c.JSON(200, res)
			}
		} else {
			return c.JSON(200, &serializer.Response{
				Status: 40001,
				Msg:    "参数错误",
				Error:  fmt.Sprint(err),
			})
		}
	}
}

//删除
func Delete(service _interface.DeleteInterface) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		if err := c.Bind(&service); err == nil {
			if err := service.Delete(); err != nil {
				return c.JSON(200, err)
			} else {
				res := service.Response()
				return c.JSON(200, res)
			}
		} else {
			return c.JSON(200, &serializer.Response{
				Status: 40001,
				Msg:    "参数错误",
				Error:  fmt.Sprint(err),
			})
		}
	}

}
