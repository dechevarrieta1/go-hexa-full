package ginmiddlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		duration := time.Since(startTime)
		statusCode := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		ipAdress := c.ClientIP()

		c.Writer.WriteString(fmt.Sprintf(
			"\n[LOG] %s | %d | %s | %s | %s\n",
			duration, statusCode, method, path, ipAdress,
		))
	}
}
