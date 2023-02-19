package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"todo_list/api"
)

func NewRouter() *gin.Engine {
	engine := gin.Default()
	store := cookie.NewStore([]byte("store-secret"))
	engine.Use(sessions.Sessions("my-session", store))
	v1Group := engine.Group("/api/v1")
	{
		v1Group.POST("user/register", api.UserRegister)
		v1Group.POST("user/login", api.UserLogin)
	}
	return engine
}
