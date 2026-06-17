package youtube

import (
	"fmt"

	"github.com/Beesonn/dlkitgo-api/src/client"
	"github.com/gin-gonic/gin"
)

func YouTubeInfo(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(400, gin.H{"error": "url parameter is required"})
		return
	}

	info, err := client.Client.Youtube.GetInfo(url)
	if err != nil {
		fmt.Println("error:", err)
		c.JSON(500, gin.H{"error": "Something went wrong please contact owner"})
		return
	}
	c.JSON(200, info)
}