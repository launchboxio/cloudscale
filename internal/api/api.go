package api

import (
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Options struct {
}

type Api struct {
	srv           *gin.Engine
	svc           *Service
	snapshotCache cache.SnapshotCache
}

func New(svc *Service, channel chan struct{}, snapshotCache cache.SnapshotCache) *Api {
	r := gin.Default()

	r.Use(Logger)

	api := &Api{srv: r, svc: svc, snapshotCache: snapshotCache}
	api.registerRoutes()

	return api
}

func (a *Api) registerRoutes() {
	a.srv.Static("/ui", "./ui/build")
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
