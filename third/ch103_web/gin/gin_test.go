package gin

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestFirst(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run(":28080")
}
