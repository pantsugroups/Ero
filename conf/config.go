package conf

const (
	// 上传路径配置
	StaticPath string = "/"

	// 数据库配置
	DataBaseName     string = "Ero"
	DataBasePassword string = "bakabie"
	DataBaseUser     string = "root"
	DttaBasePort     string = "3306"

	// 邮箱配置
	SMTPHOST     = ""
	SMTPUSERNAME = ""
	SMTPPASSWORD = ""

	JWTSecret = ""
)

func ParseDataBaseConfigure() string {
	return ""
}
