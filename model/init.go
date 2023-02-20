package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var DB *gorm.DB

// InitDatabase 为什么不需要手动的关闭数据库连接：https://github.com/go-gorm/gorm/issues/3834
func InitDatabase(info string) {
	db, err := gorm.Open("mysql", info)
	if err != nil {
		fmt.Println("打开数据库失败", err)
		panic(err)
	}
	db.LogMode(gin.Mode() == gin.DebugMode)
	db.SingularTable(true)                       //设置数据库表名不加s
	db.DB().SetMaxIdleConns(10)                  //设置连接池大小
	db.DB().SetMaxOpenConns(20)                  //最大连接数
	db.DB().SetConnMaxLifetime(30 * time.Second) //设置连接最大生存时间
	DB = db
	fmt.Println("数据库连接后成功！")
	migration()
}
