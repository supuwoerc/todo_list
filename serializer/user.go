package serializer

import "todo_list/model"

type UserResponse struct {
	ID       uint   `json:"id" form:"id"`
	UserName string `json:"user_name" form:"user_name"`
	CreateAt int64  `json:"create_at" form:"create_at"`
}

// BuildUser 将用户数据序列化
func BuildUser(user model.User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
}
