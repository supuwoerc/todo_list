package main

import (
	"fmt"
	"todo_list/conf"
	"todo_list/model"
	"todo_list/routes"
)

func main() {
	conf.Init()
	defer func() {
		model.DB.Close()
		fmt.Println("数据库连接已关闭")
	}()
	gin := routes.NewRouter()
	_ = gin.Run(conf.HttpPort)
}
