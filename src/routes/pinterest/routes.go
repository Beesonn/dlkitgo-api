package pinterest

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	r.GET("/pinterest/stream", PinStream)
}