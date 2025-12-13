package main

import (
	"github.com/Beesonn/dlkitgo-api/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupAllRoutes(r)
	r.Run(":8080")
}
