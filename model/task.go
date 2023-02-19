package model

import "github.com/jinzhu/gorm"

type TaskStatus int

const (
	UNFINISHED TaskStatus = 0 //未完成
	FINISHED   TaskStatus = 1 //已完成
)

type Task struct {
	gorm.Model
	User      User       `gorm:"ForeignKey:Uid"`
	Uid       uint       `gorm:"not null"`
	Title     string     `gorm:"index;not null"`
	Status    TaskStatus `gorm:"default:'0'"` //0未完成 1已完成
	Content   string     `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64
}
