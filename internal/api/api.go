package api

import (
	"errors"
	"fmt"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Options struct {
}

type Api struct {
	srv           *gin.Engine
	svc           *Service
	snapshotCache cache.SnapshotCache
}

func New(svc *Service, snapshotCache cache.SnapshotCache) *Api {
	r := gin.Default()

	r.Use(Logger)

	api := &Api{srv: r, svc: svc, snapshotCache: snapshotCache}
	api.registerRoutes()

	return api
}

func (a *Api) registerRoutes() {
	a.srv.Use(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/ui") {
			// Check if the file exists. If it doesn't, change the static
			// path to ./ui/build/index.html
			fileName := strings.TrimPrefix(c.Request.URL.Path, "/ui")
			if _, err := os.Stat(filepath.Join("./ui/build", fileName)); errors.Is(err, os.ErrNotExist) {
				// File not found, change file path
				fmt.Println("Setting a new filepath")
				c.Request.URL.Path = "/ui/index.html"
			}
		}
		c.Next()
	}).Static("/ui", "./ui/build")
	a.srv.GET("/healthy", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	a.srv.GET("/snapshot", func(c *gin.Context) {
		contents, err := a.snapshotCache.GetSnapshot("test-id")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, contents)
	})
	registerListenerRoutes(a.srv, a.svc)
	registerCertificateRoutes(a.srv, a.svc)
	registerTargetGroupRoutes(a.srv, a.svc)
}

func (a *Api) Run(bindAddress string) error {
	return a.srv.Run(bindAddress)
}
