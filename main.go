package main

import (
	"todo_list/conf"
	"todo_list/routes"
)

func main() {
	conf.Init()
	gin := routes.NewRouter()
	_ = gin.Run(conf.HttpPort)
}
