package router

import (
	"github.com/gin-gonic/gin"
	"video/api/drama"
	"video/api/video"
)

func InitDramaRouter(Router *gin.RouterGroup) {
	DramaRouter := Router.Group("drama")
	{
		DramaRouter.GET("", drama.List)
		//DramaRouter.GET("/:id", drama.Detail)
		DramaRouter.POST("", drama.Create)
		DramaRouter.PUT("/:id", drama.Update)
		//DramaRouter.DELETE("/:id", drama.Delete)

		// 剧集
		DramaRouter.POST("/:id/episode", video.Create)
		DramaRouter.PUT("/:id/episode/:videoId", video.Update)
		DramaRouter.DELETE("/:id/episode/:videoId", video.Delete)
	}
}
