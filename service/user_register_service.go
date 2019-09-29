package service
import (
	model "eroauz/models"
	"eroauz/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" null:"false"`
	UserName        string `form:"user_name" json:"user_name" null:"false"`
	Mail            string `form:"mail" json:"mail" null:"false"`
	Password        string `form:"password" json:"password" null:"false"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" null:"false"`
}

// Valid 验证表单
func (service *UserRegisterService) Valid() *serializer.Response {
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
func (service *UserRegisterService) SendMail() *serializer.Response{
	return nil
}
// Register 用户注册
func (service *UserRegisterService) Register() (model.User, *serializer.Response) {
	user := model.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Status:   model.Inactive,
	}

	// 表单验证
	if err := service.Valid(); err != nil {
		return user, err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return user, &serializer.Response{
			Status: 40002,
			Msg:    "密码加密失败",
		}
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return user, &serializer.Response{
			Status: 40002,
			Msg:    "注册失败",
		}
	}

	return user, nil
}