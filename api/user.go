package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_list/serializer"
	"todo_list/service"
)

func UserRegister(c *gin.Context) {
	var registerDTO service.RegisterDTO
	if err := c.ShouldBind(&registerDTO); err == nil {
		res := registerDTO.UserRegister()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Status:  http.StatusBadRequest,
			Data:    nil,
			Message: "结构体不合法",
			Error:   err.Error(),
		})
	}
}

func UserLogin(c *gin.Context) {
	var loginDTO service.LoginDTO
	if err := c.ShouldBind(&loginDTO); err == nil {
		res := loginDTO.UserLogin()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Status:  http.StatusBadRequest,
			Data:    nil,
			Message: "结构体不合法",
			Error:   err.Error(),
		})
	}
}
