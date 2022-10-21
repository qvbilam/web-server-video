package router

import (
	"github.com/gin-gonic/gin"
	"video/api/drama"
)

func InitDramaRouter(Router *gin.RouterGroup) {
	DramaRouter := Router.Group("drama")
	{
		DramaRouter.GET("", drama.List)
		//DramaRouter.GET("/:id", drama.Detail)
		//DramaRouter.POST("", drama.Create)
		//DramaRouter.PUT("/:id", drama.Update)
		//DramaRouter.DELETE("/:id", drama.Delete)
	}
}
