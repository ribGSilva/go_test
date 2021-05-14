package endpoits

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHello(c *gin.Context) {
	name := c.Param("name")
	message := "Hello " + name
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
