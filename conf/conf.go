package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
	"todo_list/model"
)

var (
	AppMode  string
	HttpPort string
	TokenKey string

	RedisAddr   string
	RedisPw     string
	RedisDbName string
	Db          string
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassword  string
	DbName      string
)

// 读取配置文件
func loadConfigFile() *ini.File {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误", err)
	}
	return file
}

// 读取项目配置
func loadServe(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").Value()
	HttpPort = file.Section("service").Key("HttpPort").Value()
	TokenKey = file.Section("service").Key("TokenKey").Value()
}

// 读取redis配置
func loadRedis(file *ini.File) {
	RedisAddr = file.Section("redis").Key("RedisAddr").Value()
	RedisPw = file.Section("redis").Key("RedisPw").Value()
	RedisDbName = file.Section("redis").Key("RedisDbName").Value()
}

// 读取mysql配置
func loadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").Value()
	DbHost = file.Section("mysql").Key("DbHost").Value()
	DbPort = file.Section("mysql").Key("DbPort").Value()
	DbUser = file.Section("mysql").Key("DbUser").Value()
	DbPassword = file.Section("mysql").Key("DbPassword").Value()
	DbName = file.Section("mysql").Key("DbName").Value()
}

// 初始化全局但数据库连接
func initMysqlConnection() {
	info := strings.Join([]string{
		DbUser,
		":",
		DbPassword,
		"@tcp(",
		DbHost,
		":",
		DbPort,
		")/",
		DbName,
		"?charset=utf8&parseTime=true",
	}, "")
	model.InitDatabase(info)
}
func Init() {
	file := loadConfigFile()
	loadServe(file)
	loadMysql(file)
	loadRedis(file)
	initMysqlConnection()
}
