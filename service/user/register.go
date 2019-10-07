package user

import (
	"eroauz/conf"
	model "eroauz/models"
	"eroauz/serializer"
	"eroauz/utils"
	"fmt"
)

// UserRegisterService 管理用户注册服务
type RegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" null:"false"`
	UserName        string `form:"username" json:"user_name" null:"false"`
	Mail            string `form:"mail" json:"mail" null:"false"`
	Password        string `form:"password" json:"password" null:"false"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" null:"false"`
}

// Valid 验证表单
func (service *RegisterService) Valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Status: 40001,
			Msg:    "两次输入的密码不相同",
		}
	}

	count := 0
	model.DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Status: 40001,
			Msg:    "昵称被占用",
		}
	}

	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Status: 40001,
			Msg:    "用户名已经注册",
		}
	}

	return nil
}

// SendMail 发送电子邮件
func (service *RegisterService) SendMail() *serializer.Response {
	//auth := smtp.PlainAuth(
	//	"",
	//	conf.SMTPUSERNAME,
	//	conf.SMTPPASSWORD,
	//	conf.SMTPHOST,
	//)
	hash := utils.Generate(service.UserName)
	token := utils.RandStringRunes(16)
	s := "您的验证地址如下：https://%s/api/v1/user/register?hash=%s&token=%s&user=%s"
	body := fmt.Sprintf(s, conf.BackEndHost, hash, token, service.UserName)
	if err := utils.SendToMail(conf.SMTPUSERNAME, conf.SMTPPASSWORD, conf.SMTPHOST, service.Mail, "Ero 注册邮件", body, "html"); err != nil {
		return &serializer.Response{
			Status: 500,
			Msg:    "邮件发送失败",
			Error:  err.Error(),
		}
	}
	//if err := smtp.SendMail(
	//	conf.SMTPHOST,
	//	auth,
	//	conf.SMTPSENDER,
	//	[]string{service.Mail},
	//	[]byte(fmt.Sprintf(body,conf.BackEndHost,hash,token,service.UserName)),
	//);err != nil{
	//	return  &serializer.Response{
	//		Status: 500,
	//		Msg:    "邮件发送失败",
	//		Error:err.Error(),
	//	}
	//}

	return nil
}

// EroAPI godoc
// @Summary 用户注册
// @Description
// @Tags user
// @Accept html
// @Produce json
// @Success 200 {object} serializer.UserResponse
// @Failure 500 {object} serializer.Response
// @Param username formData string true "用户名"
// @Param nickname formData string true "昵称（可更改）"
// @Param password formData string true "密码"
// @Param mail formData string true "邮箱"
// @Param password_confirm formData string true "重复一遍密码用于验证"
// @Router /api/v1/user/register [post]
func (service *RegisterService) Register() (model.User, *serializer.Response) {
	user := model.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Status:   model.Inactive,
		Point:    250,
		Avatar:   conf.DefaultAvatar,
	}

	// 表单验证
	if err := service.Valid(); err != nil {
		return user, err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return user, &serializer.Response{
			Status: 500,
			Msg:    "密码加密失败",
			Error:  err.Error(),
		}
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return user, &serializer.Response{
			Status: 500,
			Msg:    "注册失败",
			Error:  err.Error(),
		}
	}
	if err := service.SendMail(); err != nil {
		return user, err
	}
	return user, nil
}
