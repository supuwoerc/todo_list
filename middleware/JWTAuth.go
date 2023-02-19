package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"todo_list/conf"
	"todo_list/pkg/e"
	"todo_list/pkg/utils"
	"todo_list/serializer"
)

func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader(conf.TokenKey)
		if token == "" {
			c.Abort()
			c.JSON(http.StatusForbidden, serializer.Response{
				Status:  e.ErrorAuth,
				Data:    nil,
				Message: e.GetMessage(e.ErrorAuth),
				Error:   "",
			})
			return
		}
		parseToken, err := utils.ParseToken(token)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusForbidden, serializer.Response{
				Status:  e.ErrorAuthCheckTokenFail,
				Data:    nil,
				Message: e.GetMessage(e.ErrorAuthCheckTokenFail),
				Error:   err.Error(),
			})
		} else {
			if time.Now().Unix() > parseToken.ExpiresAt {
				c.Abort()
				c.JSON(http.StatusForbidden, serializer.Response{
					Status:  e.ErrorAuthCheckTokenTimeout,
					Data:    nil,
					Message: e.GetMessage(e.ErrorAuthCheckTokenTimeout),
					Error:   "",
				})
			}
		}
		c.Next()
	}
}
