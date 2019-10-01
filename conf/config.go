package conf

const (
	// 上传路径配置
	Static_Path string = "/"

	// 数据库配置
	DataBase_Name     string = "Ero"
	DataBase_Password string = "bakabie"
	DataBase_User     string = "root"
	DttaBase_Port     string = "3306"

	// 邮箱配置
	SMTP_HOST     = ""
	SMTP_USERNAME = ""
	SMTP_PASSWORD = ""
)

func ParseDataBaseConfigure() string {
	return ""
}
