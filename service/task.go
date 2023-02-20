package service

import (
	"github.com/jinzhu/gorm"
	"time"
	"todo_list/model"
	"todo_list/pkg/e"
	"todo_list/serializer"
)

type TaskDTO struct {
	Title   string `json:"title" form:"title" binding:"required,min=1,max=10"`
	Content string `json:"content" form:"content" binding:"required,min=1"`
}
type TaskDetailDTO struct {
	ID uint `uri:"tid" binding:"required"`
}

// TaskCreate 创建一个task
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

// TaskDetail 查询task详情
func (taskDetailDTO *TaskDetailDTO) TaskDetail(uid uint) serializer.Response {
	var task model.Task
	if err := model.DB.Model(&task).Where("id = ? and uid = ?", taskDetailDTO.ID, uid).First(&task).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status:  e.ErrorTaskNotFound,
				Data:    nil,
				Message: e.GetMessage(e.ErrorTaskNotFound),
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
	return serializer.Response{
		Status:  e.SUCCESS,
		Data:    serializer.BuildTaskDetail(task),
		Message: e.GetMessage(e.SUCCESS),
		Error:   "",
	}
}
