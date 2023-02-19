package service

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"todo_list/model"
	"todo_list/pkg/utils"
	"todo_list/serializer"
)

type RegisterDTO struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=12"`
}

// LoginDTO 登陆的dto和注册的基本一致，但是后续可能需要扩展验证码等功能，所以此处做一个区分
type LoginDTO struct {
	RegisterDTO
}

// UserRegister 用户注册
func (registerDTO *RegisterDTO) UserRegister() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", registerDTO.UserName).First(&user).Count(&count)
	if count != 0 {
		return serializer.Response{
			Status:  http.StatusBadRequest,
			Data:    nil,
			Message: "用户名重复，注册失败",
			Error:   "",
		}
	}
	user.UserName = registerDTO.UserName
	if err := user.SetPassword(registerDTO.Password); err != nil {
		return serializer.Response{
			Status:  http.StatusBadRequest,
			Data:    nil,
			Message: "用户密码加密失败",
			Error:   err.Error(),
		}
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status:  http.StatusBadRequest,
			Data:    nil,
			Message: "数据库创建用户失败",
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  http.StatusOK,
		Data:    serializer.BuildUser(user),
		Message: "注册成功",
		Error:   "",
	}
}

// UserLogin 用户登陆
func (loginDTO *LoginDTO) UserLogin() serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("user_name=?", loginDTO.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status:  http.StatusBadRequest,
				Data:    nil,
				Message: "账户不存在",
				Error:   err.Error(),
			}
		} else {
			return serializer.Response{
				Status:  http.StatusBadRequest,
				Data:    nil,
				Message: "数据库发生错误",
				Error:   err.Error(),
			}
		}
	}
	checkResult := user.CheckPassword(loginDTO.Password)
	if !checkResult {
		return serializer.Response{
			Status:  http.StatusBadRequest,
			Data:    nil,
			Message: "密码错误",
			Error:   "",
		}
	}
	token, err := utils.GenerateToken(user.ID, user.UserName)
	if err != nil {
		return serializer.Response{
			Status:  http.StatusBadRequest,
			Data:    nil,
			Message: "签发token失败",
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  http.StatusOK,
		Data:    serializer.BuildLoginResponse(serializer.BuildUser(user), token),
		Message: "登陆成功",
		Error:   "",
	}
}
