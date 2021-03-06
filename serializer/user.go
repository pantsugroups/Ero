package serializer

import model "eroauz/models"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Mail      string `json:"mail"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
	Point     int    `json:"point"`
	Bio       string `json:"bio"`
	Hito      string `json:"hito"`
	Website   string `json:"website"`
}

// UserResponse 单个用户序列化
type UserResponse struct {
	Response
	Data User `json:"data"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Avatar:    user.Avatar,
		Mail:      user.Mail,
		CreatedAt: user.CreatedAt.Unix(),
		Point:     user.Point,
		Bio:       user.Bio,
		Hito:      user.Hito,
		Website:   user.Hito,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) UserResponse {
	return UserResponse{
		Data: BuildUser(user),
	}
}
