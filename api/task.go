package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_list/conf"
	"todo_list/pkg/e"
	"todo_list/pkg/utils"
	"todo_list/serializer"
	"todo_list/service"
)

// TaskCreate 新建待办
func TaskCreate(c *gin.Context) {
	var taskDTO service.TaskDTO
	claims, _ := utils.ParseToken(c.GetHeader(conf.TokenKey))
	if err := c.ShouldBind(&taskDTO); err == nil {
		res := taskDTO.TaskCreate(claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Status:  e.InvalidParams,
			Data:    nil,
			Message: e.GetMessage(e.InvalidParams),
			Error:   err.Error(),
		})
	}
}

// TaskDetail 待办详情
func TaskDetail(c *gin.Context) {
	var taskDetailDTO service.TaskDetailDTO
	claims, _ := utils.ParseToken(c.GetHeader(conf.TokenKey))
	if err := c.ShouldBindUri(&taskDetailDTO); err == nil {
		res := taskDetailDTO.TaskDetail(claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Status:  e.InvalidParams,
			Data:    nil,
			Message: e.GetMessage(e.InvalidParams),
			Error:   err.Error(),
		})
	}
}
func TaskList(c *gin.Context) {
	var taskDTO service.TasksDTO
	claims, _ := utils.ParseToken(c.GetHeader(conf.TokenKey))
	if err := c.ShouldBindQuery(&taskDTO); err == nil {
		res := taskDTO.TaskList(claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Status:  e.InvalidParams,
			Data:    nil,
			Message: e.GetMessage(e.InvalidParams),
			Error:   err.Error(),
		})
	}
}
func TaskUpdate(c *gin.Context) {
	var taskUpdateDTO service.TaskUpdateDTO
	claims, _ := utils.ParseToken(c.GetHeader(conf.TokenKey))
	if err := c.ShouldBind(&taskUpdateDTO); err == nil {
		res := taskUpdateDTO.TaskUpdate(claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Status:  e.InvalidParams,
			Data:    nil,
			Message: e.GetMessage(e.InvalidParams),
			Error:   err.Error(),
		})
	}
}
func TaskDelete(c *gin.Context) {
	var taskDeleteDTO service.TaskDeleteDTO
	claims, _ := utils.ParseToken(c.GetHeader(conf.TokenKey))
	if err := c.ShouldBindUri(&taskDeleteDTO); err == nil {
		res := taskDeleteDTO.TaskDelete(claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Status:  e.InvalidParams,
			Data:    nil,
			Message: e.GetMessage(e.InvalidParams),
			Error:   err.Error(),
		})
	}
}
