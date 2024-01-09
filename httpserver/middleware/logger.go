package middleware

import (
	"sbm-itb/httpserver/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		param := map[string]interface{}{
			"status_code": c.Writer.Status(),
			"method":      c.Request.Method,
			"latency":     time.Since(start),
			"path":        path,
		}

		log.Info(param)

		if len(c.Errors) == 0 {
			log.Info(param)
		} else {
			errList := []error{}
			for _, err := range c.Errors {
				errList = append(errList, err)
			}

			if len(errList) > 0 {
				param["errors"] = errList
				log.Info(param)
			}
		}

		c.Next()
	}
}
