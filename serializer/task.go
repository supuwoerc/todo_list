package serializer

import "todo_list/model"

type TaskDetailResponse struct {
	id        uint             `json:"id"`
	Uid       uint             `json:"uid"`
	Title     string           `json:"title"`
	Status    model.TaskStatus `json:"status"` //0未完成 1已完成
	Content   string           `json:"content"`
	StartTime int64            `json:"start_time"`
	EndTime   int64            `json:"end_time"`
}

// BuildTaskDetail 将task数据序列化
func BuildTaskDetail(task model.Task) TaskDetailResponse {
	return TaskDetailResponse{
		id:        task.ID,
		Uid:       task.Uid,
		Title:     task.Title,
		Status:    task.Status,
		Content:   task.Content,
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
}

// BuildTasks 将tasks数据序列化
func BuildTasks(tasks []model.Task, total uint) DataList[TaskDetailResponse] {
	var list []TaskDetailResponse
	for _, v := range tasks {
		list = append(list, BuildTaskDetail(v))
	}
	return DataList[TaskDetailResponse]{
		List:  list,
		Total: total,
	}
}
