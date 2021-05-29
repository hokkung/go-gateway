package router

import (
	"net/http"

	"github.com/hokkung/go-gateway/server/auth"
	"github.com/hokkung/go-gateway/server/utils"

	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	proxy := engine.Group("/_api")
	proxy.Use(auth.JWT())
	proxy.GET("/ping", pingResponse())
}

func pingResponse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ua := ctx.Query("ua")
		_, err := utils.IsAllowedUserAgent(ua)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			return
		}

		ctx.String(http.StatusOK, ua)
	}
}
