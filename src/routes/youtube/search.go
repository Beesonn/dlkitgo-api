package youtube

import (
	"fmt"
	"strconv"

	"github.com/Beesonn/dlkitgo-api/src/client"
	"github.com/gin-gonic/gin"
)

func YouTubeSearch(c *gin.Context) {
	q := c.Query("q")
	limStr := c.Query("lim")
	if q == "" {
		c.JSON(400, gin.H{"error": "q parameter is required"})
		return
	}

	lim := 0
	if limStr != "" {
		if val, err := strconv.Atoi(limStr); err == nil {
			lim = val
		}
	}

	search, err := client.Client.Youtube.Search(q, lim)
	if err != nil {
		fmt.Println("error:", err)
		c.JSON(500, gin.H{"error": "Something went wrong please contact owner"})
		return
	}
	c.JSON(200, search)
}