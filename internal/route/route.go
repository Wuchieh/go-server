package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route
//
//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
func Route(r gin.IRouter) {
	r.GET("/", helloWorld)
}

// helloWorld
//
//	@Id			helloWorld
//	@Summary	hello world
//	@tags		example
//	@Produce	plain
//	@Success	200	{string}	string
//	@Router		/ [get]
func helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
