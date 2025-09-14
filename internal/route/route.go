package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Route(r gin.IRouter) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
}
