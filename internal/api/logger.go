package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

var timeFormat = "02/Jan/2006:15:04:05 -0700"

func Logger(c *gin.Context) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	path := c.Request.URL.Path
	start := time.Now()
	c.Next()
	stop := time.Since(start)
	latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
	status := c.Writer.Status()
	clientIp := c.ClientIP()
	clientUserAgent := c.Request.UserAgent()

	dataLength := c.Writer.Size()
	if dataLength < 0 {
		dataLength = 0
	}

	entry := log.WithFields(log.Fields{
		"status":     status,
		"latency":    latency,
		"client_ip":  clientIp,
		"user_agent": clientUserAgent,
		"path":       path,
		"method":     c.Request.Method,
	})

	if len(c.Errors) > 0 {
		entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
	} else {
		msg := fmt.Sprintf("%s - %s [%s] \"%s %s\" %d %d \"%s\" \"%s\" (%dms)", clientIp, hostname, time.Now().Format(timeFormat), c.Request.Method, path, status, dataLength, "", clientUserAgent, latency)
		if status >= http.StatusInternalServerError {
			entry.Error(msg)
		} else if status >= http.StatusBadRequest {
			entry.Warn(msg)
		} else {
			entry.Info(msg)
		}
	}
}
