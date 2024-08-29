package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	//Initialize router
	r := gin.Default()
	// Initilize routes
	initializeRoutes(r)
	// listen and serve on 0.0.0.0:8080
	r.Run()
}
