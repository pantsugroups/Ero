package conf

import "fmt"

const (
	WebPort = ":8000"
	// 后端地址，如果不填写验证邮件的URL可能会出问题
	BackEndHost = "api.ero.ink"
	// 上传路径配置
	StaticPath string = ""

	// 数据库配置
	DataBaseName     string = "ero"
	DataBasePassword string = "bakabie"
	DataBaseUser     string = "root"
	//DataBasePort     string = "3306"

	// 邮箱配置
	SMTPHOST     = ""
	SMTPUSERNAME = ""
	SMTPPASSWORD = ""

	Secret = "secret"

	DefaultAvatar = "https://api.ero.ink/img/default_avatar.jpg"
	DefaultCover  = "https://api.ero.ink/img/default_cover.jpg"
)

func ParseDataBaseConfigure() string {
	s := "%s:%s@/%s?charset=utf8&parseTime=True&loc=Local"
	return fmt.Sprintf(s, DataBaseUser, DataBasePassword, DataBaseName)
}
