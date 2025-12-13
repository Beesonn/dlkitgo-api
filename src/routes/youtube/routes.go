package youtube

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	r.GET("/youtube/stream", YouTubeStream)
}
