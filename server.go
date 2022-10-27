package main

import (
	"github.com/gin-gonic/gin"
	controller "github.com/raghavkx/go-gin/controllers"
	"github.com/raghavkx/go-gin/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})
	server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(
			200, videoController.FindAll(),
		)

	})
	server.POST("/posts", func(ctx *gin.Context) {
		ctx.JSON(
			200,
			videoController.Save(ctx),
		)
	})
	server.Run(":8080")
}
