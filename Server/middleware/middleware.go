package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckMethod(c *gin.Context) {
	start := time.Now()
	c.Next()
	fmt.Printf("[Middleware] Request URL: %s | Method: %s | Time: %s | Status: %d\n",
		c.Request.URL.Path, c.Request.Method, start.Format(time.ANSIC), c.Writer.Status())
}
