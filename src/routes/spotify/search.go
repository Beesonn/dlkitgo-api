package spotify

import (
	"fmt"

	"github.com/Beesonn/dlkitgo-api/src/client"
	"github.com/gin-gonic/gin"
)

func SpotifySearch(c *gin.Context) {
	q := c.Query("q")
	if q == "" {
		c.JSON(400, gin.H{"error": "q parameter is required"})
		return
	}

	stream, err := client.Client.Spotify.Search(q)
	if err != nil {
		fmt.Println("error:", err)
		c.JSON(500, gin.H{"error": "Something went wrong please contact owner"})
		return
	}
	c.JSON(200, stream)
}
