package spotify

import (
	"fmt"

	"github.com/Beesonn/dlkitgo-api/src/client"
	"github.com/gin-gonic/gin"
)

func SpotifyStream(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(400, gin.H{"error": "url parameter is required"})
		return
	}

	stream, err := client.Client.Spotify.Stream(url)
	if err != nil {
		fmt.Println("error:", err)
		c.JSON(500, gin.H{"error": "Something went wrong please contact owner"})
		return
	}
	if len(stream.Source) == 0 {
		fmt.Println("No stream sources found")
		c.JSON(404, gin.H{"error": "Something went wrong please contact owner"})
		return
	}
	c.JSON(200, stream)
}
