package youtube

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	r.GET("/youtube/stream", YouTubeStream)
	r.GET("/youtube/search", YouTubeSearch)
	r.GET("/youtube/info", YouTubeInfo)
}