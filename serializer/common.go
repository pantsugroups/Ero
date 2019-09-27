package serializer

import "eroauz/models"

// Response 团队基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

type TokenResponse struct {
	Response
	Token string `json:"token"`
	Data  User    `json:"data"`
}
// 创建用户第一次注册时包含Token的Response
func BuildTokenResponse(u models.User,token string) *TokenResponse{
	return &TokenResponse{
		Token:token,
		Data:BuildUser(u),
	}
}