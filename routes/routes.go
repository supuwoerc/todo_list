package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"todo_list/api"
	"todo_list/middleware"
)

func NewRouter() *gin.Engine {
	engine := gin.Default()
	store := cookie.NewStore([]byte("store-secret"))
	engine.Use(sessions.Sessions("my-session", store))
	v1Group := engine.Group("/api/v1")
	{
		v1Group.POST("user/register", api.UserRegister)
		v1Group.POST("user/login", api.UserLogin)
		v1AuthGroup := v1Group.Group("/") //需要鉴权的路由
		v1AuthGroup.Use(middleware.JWTAuth())
		{
			v1AuthGroup.POST("task", api.TaskCreate)
			v1AuthGroup.GET("task/:tid", api.TaskDetail)
			v1AuthGroup.GET("tasks", api.TaskList)
		}
	}
	return engine
}
