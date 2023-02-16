package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandlerFn() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	}
}
