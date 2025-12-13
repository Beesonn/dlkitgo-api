package spotify

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	r.GET("/spotify/stream", SpotifyStream)
	r.GET("/spotify/search", SpotifySearch)
	r.GET("/spotify/info", SpotifyInfo)
}
