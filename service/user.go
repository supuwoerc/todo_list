package service

import (
	"github.com/jinzhu/gorm"
	"todo_list/model"
	"todo_list/pkg/e"
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
			Status:  e.ErrorExistUser,
			Data:    nil,
			Message: e.GetMessage(e.ErrorExistUser),
			Error:   "",
		}
	}
	user.UserName = registerDTO.UserName
	if err := user.SetPassword(registerDTO.Password); err != nil {
		return serializer.Response{
			Status:  e.ErrorFailEncryption,
			Data:    nil,
			Message: e.GetMessage(e.ErrorFailEncryption),
			Error:   err.Error(),
		}
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status:  e.ErrorDatabase,
			Data:    nil,
			Message: e.GetMessage(e.ErrorDatabase),
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  e.SUCCESS,
		Data:    serializer.BuildUser(user),
		Message: e.GetMessage(e.SUCCESS),
		Error:   "",
	}
}

// UserLogin 用户登陆
func (loginDTO *LoginDTO) UserLogin() serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("user_name=?", loginDTO.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status:  e.ErrorNotExistUser,
				Data:    nil,
				Message: e.GetMessage(e.ErrorNotExistUser),
				Error:   err.Error(),
			}
		} else {
			return serializer.Response{
				Status:  e.ErrorDatabase,
				Data:    nil,
				Message: e.GetMessage(e.ErrorDatabase),
				Error:   err.Error(),
			}
		}
	}
	checkResult := user.CheckPassword(loginDTO.Password)
	if !checkResult {
		return serializer.Response{
			Status:  e.ErrorNotCompare,
			Data:    nil,
			Message: e.GetMessage(e.ErrorNotCompare),
			Error:   "",
		}
	}
	token, err := utils.GenerateToken(user.ID, user.UserName)
	if err != nil {
		return serializer.Response{
			Status:  e.ErrorAuthToken,
			Data:    nil,
			Message: e.GetMessage(e.ErrorAuthToken),
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  e.SUCCESS,
		Data:    serializer.BuildLoginResponse(serializer.BuildUser(user), token),
		Message: e.GetMessage(e.SUCCESS),
		Error:   "",
	}
}
