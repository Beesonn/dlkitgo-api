package instagram

import (
	"fmt"

	"github.com/Beesonn/dlkitgo-api/src/client"
	"github.com/gin-gonic/gin"
)

func InstaInfo(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(400, gin.H{"error": "url parameter is required"})
		return
	}

	stream, err := client.Client.Instagram.GetInfo(url)
	if err != nil {
		fmt.Println("error:", err)
		c.JSON(500, gin.H{"error": "Something went wrong please contact owner"})
		return
	}
	c.JSON(200, stream)
}
