package routes

import (
	"github.com/Beesonn/dlkitgo-api/src/routes/instagram"
	"github.com/Beesonn/dlkitgo-api/src/routes/pinterest"
	"github.com/Beesonn/dlkitgo-api/src/routes/spotify"
	"github.com/Beesonn/dlkitgo-api/src/routes/youtube"
	"github.com/gin-gonic/gin"
)

func SetupAllRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello baby I'm working fine",
			"available_endpoints": gin.H{
				"instagram": []string{
					"GET /instagram/stream?url=instagram_url",
					"GET /instagram/info?url=instagram_url",
				},
				"youtube": []string{
					"GET /youtube/stream?url=youtube_url",
				},
				"spotify": []string{
					"GET /spotify/search?q=song_name",
					"GET /spotify/info?url=spotify_url",
					"GET /spotify/stream?url=spotify_url",
				},
				"pinterest": []string{
					"GET /pinterest/stream?url=pinterest_url",
				},
			},
		})
	})
	spotify.SetupRoutes(r)
	youtube.SetupRoutes(r)
	instagram.SetupRoutes(r)
	pinterest.SetupRoutes(r)
}
