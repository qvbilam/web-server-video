package router

import (
	"github.com/gin-gonic/gin"
	"video/api/episodes"
)

func InitEpisodesRouter(Router *gin.RouterGroup) {
	VideoRouter := Router.Group("video/episodes")
	{
		VideoRouter.GET("", episodes.List)
		VideoRouter.GET("/:id", episodes.Detail)
		VideoRouter.POST("", episodes.Create)
		VideoRouter.PUT("/:id", episodes.Update)
		VideoRouter.DELETE("/:id", episodes.Delete)
	}
}
