package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Options struct {
}

type Api struct {
	srv *gin.Engine
	svc *Service
}

func New(svc *Service, channel chan struct{}) *Api {
	r := gin.Default()

	r.Use(Logger)

	api := &Api{srv: r, svc: svc}
	api.registerRoutes()

	return api
}

func (a *Api) registerRoutes() {
	a.srv.GET("/healthy", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	registerListenerRoutes(a.srv, a.svc)
	registerCertificateRoutes(a.srv, a.svc)
	registerTargetGroupRoutes(a.srv, a.svc)
}

func (a *Api) Run(bindAddress string) error {
	return a.srv.Run(bindAddress)
}
