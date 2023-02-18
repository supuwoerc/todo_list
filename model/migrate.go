package model

import "fmt"

// 数据库自动迁移
func migration() {
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{}).AutoMigrate(&Task{})
	DB.Model(&Task{}).AddForeignKey("Uid", "User(id)", "CASCADE", "CASCADE")
	fmt.Println("数据库迁移成功")
}
