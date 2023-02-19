package service

import (
	"time"
	"todo_list/model"
	"todo_list/pkg/e"
	"todo_list/serializer"
)

type TaskDTO struct {
	Title   string `json:"title" form:"title" binding:"required,min=1,max=10"`
	Content string `json:"content" form:"content" binding:"required,min=1"`
}

func (taskDTO *TaskDTO) TaskCreate(uid uint) serializer.Response {
	var user model.User
	model.DB.First(&user, uid)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     taskDTO.Title,
		Status:    model.UNFINISHED,
		Content:   taskDTO.Content,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	if err := model.DB.Create(&task).Error; err == nil {
		return serializer.Response{
			Status:  e.SUCCESS,
			Data:    nil,
			Message: e.GetMessage(e.SUCCESS),
			Error:   "",
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
