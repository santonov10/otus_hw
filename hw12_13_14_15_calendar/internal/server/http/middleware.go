package internalhttp

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/logger"
)

func LoggingMiddleware(logger *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		if raw != "" {
			path = path + "?" + raw
		}

		logger.Info().Msgf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			c.ClientIP(),
			time.Now().Format(time.RFC1123),
			c.Request.Method,
			path,
			c.Request.Proto,
			c.Writer.Status(),
			time.Now().Sub(start),
			c.Request.UserAgent(),
			c.Errors.ByType(gin.ErrorTypePrivate).String(),
		)
	}
}
