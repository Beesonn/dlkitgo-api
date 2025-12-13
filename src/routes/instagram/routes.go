package instagram

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	r.GET("/instagram/stream", InstaStream)
	r.GET("/instagram/info", InstaInfo)
}
