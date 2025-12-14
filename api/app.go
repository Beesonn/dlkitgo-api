package app

import (
	"net/http"

	"github.com/Beesonn/dlkitgo-api/src/routes"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	r = gin.Default()
	routes.SetupAllRoutes(r)
}

func Handler(w http.ResponseWriter, req *http.Request) {
	r.ServeHTTP(w, req)
}
