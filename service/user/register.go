package user

import (
	"eroauz/conf"
	model "eroauz/models"
	"eroauz/serializer"
	"eroauz/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

// UserRegisterService 管理用户注册服务
type RegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" null:"false"`
	UserName        string `form:"username" json:"user_name" null:"false"`
	Mail            string `form:"mail" json:"mail" null:"false"`
	Password        string `form:"password" json:"password" null:"false"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" null:"false"`
	VerifyCode      string `json:"verify_code" form:"verify_code" null:"false"`
	VerifyCodeId    string `json:"verify_id" form:"verify_id" null:"false"`
	InviteCode      string `json:"invite_code" form:"invite_code"`
}

// Valid 验证表单
func (service *RegisterService) Valid() *serializer.Response {
	var invite model.Invite
	if conf.AllowRegister == false {
		fmt.Println(service.InviteCode)
		if service.InviteCode == "" {
			return &serializer.Response{
				Status: 500,
				Msg:    "请输入邀请码",
			}
		}
		if err := model.DB.Where(&model.Invite{Code: service.InviteCode}).First(&invite).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				return &serializer.Response{
					Status: 500,
					Msg:    "邀请码错误或者过期",
				}
			}
			return &serializer.Response{
				Status: 500,
				Msg:    "数据库错误",
			}
		}
		//if invite.Code == ""{
		//	return &serializer.Response{
		//		Status: 500,
		//		Msg:    "邀请码错误或者过期",
		//	}
		//}
		if time.Now().After(invite.TimeLimit) {
			return &serializer.Response{
				Status: 500,
				Msg:    "邀请码错误或者过期",
			}
		}

	}
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Status: 500,
			Msg:    "两次输入的密码不相同",
		}
	}
	if len(service.Password) < 6 || len(service.Password) > 16 {
		return &serializer.Response{
			Status: 500,
			Msg:    "密码格式错误，请大于6位且小于16位数",
		}
	}
	count := 0
	model.DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Status: 500,
			Msg:    "昵称被占用",
		}
	}

	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Status: 500,
			Msg:    "用户名已经注册",
		}
	}

	return nil
}

// SendMail 发送电子邮件
func (service *RegisterService) SendMail() *serializer.Response {
	hash := utils.Generate(service.UserName)
	token := utils.RandStringRunes(16)
	s := "您的验证地址如下：https://%s/api/v1/user/register?hash=%s&token=%s&user=%s"
	body := fmt.Sprintf(s, conf.BackEndHost, hash, token, service.UserName)
	if err := utils.SendToMail(conf.SMTPUSERNAME, conf.SMTPPASSWORD, conf.SMTPHOST, service.Mail, "Ero 注册邮件", body, "text"); err != nil {
		return nil
	}

	return nil
}

// EroAPI godoc
// @Summary 用户注册
// @Description 必须要先从/api/v1/verify 处获取验证码
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
// @Param verify_code formData string true "验证码"
// @Param verify_id formData string true "验证码ID"
// @Router /api/v1/user/register [post]
func (service *RegisterService) Register() (model.User, *serializer.Response) {
	if res := utils.VerifyCaptcha(service.VerifyCodeId, service.VerifyCode); res == false {
		return model.User{}, &serializer.Response{
			Status: 403,
			Msg:    "验证码错误",
		}
	}
	user := model.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Mail:     service.Mail,
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
