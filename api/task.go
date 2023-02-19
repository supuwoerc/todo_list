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
	//此处不需要处理错误，因为路由已经确保token是合法的
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
