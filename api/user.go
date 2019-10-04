package api

import (
	"eroauz/models"
	"eroauz/serializer"
	"eroauz/service/user"
	"eroauz/utils"
	"fmt"
	"github.com/labstack/echo"
)

func UserRegister(c echo.Context) (err error) {

	var service user.RegisterService
	if err := c.Bind(&service); err == nil {
		if users, err := service.Register(); err != nil {
			return c.JSON(200, err)
		} else {
			res := serializer.BuildUserResponse(users)
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
func UserLogin(c echo.Context) (err error) {
	var service user.LoginService
	if err := c.Bind(&service); err == nil {
		if users, err := service.Login(); err != nil {
			return c.JSON(200, err)
		} else {
			// 设置Session
			token, err := utils.CreateToken(users)
			if err != nil {
				return c.JSON(200, &serializer.Response{
					Status: 40003,
					Msg:    "token生成失败",
					Error:  fmt.Sprint(err)})
			}
			res := serializer.BuildTokenResponse(users, token)
			return c.JSON(200, res)
		}
	} else {
		return c.JSON(200, &serializer.Response{
			Status: 40001,
			Msg:    "参数错误",
			Error:  fmt.Sprint(err)})
	}
}
func VeruftMail(c echo.Context) error {
	token := c.QueryParam("token")
	s := c.QueryParam("user")
	hash := c.QueryParam("hash")
	if len(token) != 16 {
		return c.JSON(200, serializer.Response{
			Status: 403,
			Msg:    "验证失败",
		})
	}
	u, err := models.GetUser(s)
	if err != nil {
		return c.JSON(200, serializer.Response{
			Status: 404,
			Msg:    "用户不存在",
		})
	}
	if hash != utils.Generate(s) {
		return c.JSON(200, serializer.Response{
			Status: 403,
			Msg:    "令牌错误",
		})
	} else {

		if err := models.DB.Model(&u).Update(models.User{
			Status: models.Active,
		}).Error; err != nil {
			return c.JSON(200, serializer.Response{
				Status: 500,
				Msg:    "激活失败",
				Error:  err.Error(),
			})
		} else {
			return c.JSON(200, serializer.Response{
				Status: 0,
				Msg:    "激活成功",
			})
		}

	}
}
