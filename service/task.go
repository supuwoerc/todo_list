package service

import (
	"github.com/jinzhu/gorm"
	"strings"
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
type TasksDTO struct {
	PageSize int    `form:"page_size"`
	Page     int    `form:"page"`
	Keyword  string `form:"keyword"`
}
type TaskUpdateDTO struct {
	ID      uint   `json:"id" form:"id" binding:"required"`
	Title   string `json:"title" form:"title" binding:"required,min=1,max=10"`
	Content string `json:"content" form:"content" binding:"required,min=1"`
}
type TaskDeleteDTO struct {
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

// TaskList 查询分页Task
func (tasksDTO *TasksDTO) TaskList(uid uint) serializer.Response {
	var total uint
	var tasks []model.Task
	if tasksDTO.Page <= 0 {
		tasksDTO.Page = 1
	}
	if tasksDTO.PageSize <= 0 {
		tasksDTO.PageSize = 10
	}
	keyword := strings.Join([]string{"%", tasksDTO.Keyword, "%"}, "")
	err := model.DB.Model(model.Task{}).Preload("User").Where("uid = ? and (title like ? or content like ?)", uid, keyword, keyword).Count(&total).
		Limit(tasksDTO.PageSize).Offset((tasksDTO.Page - 1) * tasksDTO.PageSize).
		Find(&tasks).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return serializer.Response{
			Status:  e.ErrorTaskNotFound,
			Data:    nil,
			Message: e.GetMessage(e.ErrorTaskNotFound),
			Error:   err.Error(),
		}
	} else if err != nil {
		return serializer.Response{
			Status:  e.ErrorDatabase,
			Data:    nil,
			Message: e.GetMessage(e.ErrorDatabase),
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  e.SUCCESS,
		Data:    serializer.BuildTasks(tasks, total),
		Message: e.GetMessage(e.SUCCESS),
		Error:   "",
	}
}

// TaskUpdate 更新task
func (taskUpdateDTO *TaskUpdateDTO) TaskUpdate(uid uint) serializer.Response {
	var task model.Task
	if err := model.DB.Where("uid = ? and id = ?", uid, taskUpdateDTO.ID).Find(&task).Error; err != nil {
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
	task.Title = taskUpdateDTO.Title
	task.Content = taskUpdateDTO.Content
	if err := model.DB.Save(&task).Error; err != nil {
		return serializer.Response{
			Status:  e.ERROR,
			Data:    nil,
			Message: e.GetMessage(e.ERROR),
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  e.SUCCESS,
		Data:    nil,
		Message: e.GetMessage(e.SUCCESS),
		Error:   "",
	}
}

func (taskDeleteDTO *TaskDeleteDTO) TaskDelete(uid uint) serializer.Response {
	var task model.Task
	if err := model.DB.Model(&task).Where("id = ? and uid = ?", taskDeleteDTO.ID, uid).Find(&task).Error; err != nil {
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
	if err := model.DB.Delete(&task).Error; err != nil {
		return serializer.Response{
			Status:  e.ErrorDatabase,
			Data:    nil,
			Message: e.GetMessage(e.ErrorDatabase),
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  e.SUCCESS,
		Data:    true,
		Message: e.GetMessage(e.SUCCESS),
		Error:   "",
	}

}
