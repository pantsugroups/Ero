package api

import (
	"eroauz/conf"
	"eroauz/models"
	"eroauz/serializer"
	"eroauz/service/user"
	"eroauz/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"time"
)

func UserRegister(c echo.Context) (err error) {

	var service user.RegisterService
	if err := utils.Bind(&service, c); err == nil {
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
	if err := utils.Bind(&service, c); err == nil {
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

// EroAPI godoc
// @Summary 发送用户验证邮件
// @Description 必须要先从/api/v1/verify 处获取验证码
// @Tags user
// @Accept html
// @Produce json
// @Success 200 {object} serializer.Response
// @Failure 500 {object} serializer.Response
// @Param verify_code query string true "验证码"
// @Param verify_id query string true "验证码ID"
// @Router /api/v1/user/sendmail [get]
func SendMail(c echo.Context) error {
	var InviteMail models.InviteMail
	//verifyID := c.QueryParam("verify_id")
	//verifyCode := c.QueryParam("verify_code")
	//if res := utils.VerifyCaptcha(verifyID, verifyCode); res == false {
	//	return c.JSON(200, &serializer.Response{
	//		Status: 403,
	//		Msg:    "验证码错误",
	//	})
	//}
	uid := utils.GetAuthorID(c)
	u, err := models.GetUser(uid)
	if err != nil {
		return c.JSON(200, &serializer.Response{
			Status: 404,
			Msg:    "没有找到该用户",
			Error:  err.Error()})
	}
	if err := models.DB.Where(&models.InviteMail{User: u.ID}).First(&InviteMail).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return c.JSON(200, &serializer.Response{
				Status: 500,
				Msg:    "数据库错误",
				Error:  err.Error(),
			})
		} else {
			if time.Now().Before(InviteMail.TimeLimit) {
				return c.JSON(200, &serializer.Response{
					Status: 500,
					Msg:    "间隔太快",
				})
			}
		}
	}
	if u.Status == models.Inactive {
		hash := utils.Generate(u.UserName)
		token := utils.RandStringRunes(16)
		s := "您的验证地址如下：https://%s/api/v1/user/register?hash=%s&token=%s&user=%s"
		body := fmt.Sprintf(s, conf.BackEndHost, hash, token, u.UserName)
		if err := utils.SendToMail(conf.SMTPUSERNAME, conf.SMTPPASSWORD, conf.SMTPHOST, u.Mail, "Ero 注册邮件", body, "html"); err != nil {
			return c.JSON(200, &serializer.Response{
				Status: 500,
				Msg:    "邮件发送失败",
				Error:  err.Error()})
		}
		dd, _ := time.ParseDuration("1m")
		limit := time.Now().Add(dd)
		InviteMail = models.InviteMail{
			TimeLimit: limit,
			User:      u.ID,
		}

		if err := models.DB.Create(&InviteMail).Error; err != nil {
			return c.JSON(200, &serializer.Response{
				Status: 500,
				Msg:    "邮件发送失败",
				Error:  err.Error()})
		}
		return c.JSON(200, &serializer.Response{
			Status: 0,
			Msg:    "邮件发送成功"})
	} else {
		return c.JSON(200, &serializer.Response{
			Status: 0,
			Msg:    "您已经是会员，无需再次验证"})
	}

}
func VerifyMail(c echo.Context) error {
	token := c.QueryParam("token")
	s := c.QueryParam("user")
	hash := c.QueryParam("hash")
	if len(token) != 16 {
		return c.JSON(200, serializer.Response{
			Status: 403,
			Msg:    "验证失败",
		})
	}
	var u models.User
	err := models.DB.Where("user_name = ?", s).First(&u).Error
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
