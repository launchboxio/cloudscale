package api

import (
	"github.com/gin-gonic/gin"
	bolt "go.etcd.io/bbolt"
	"net/http"
)

type Options struct {
}

type Api struct {
	srv *gin.Engine
	db  *bolt.DB
}

func New(db *bolt.DB) *Api {
	r := gin.Default()

	r.Use(Logger)

	api := &Api{srv: r, db: db}
	api.registerRoutes()

	return api
}

func (a *Api) registerRoutes() {
	a.srv.GET("/healthy", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	registerListenerRoutes(a.srv, a.db)
	registerCertificateRoutes(a.srv, a.db)
	registerTargetGroupRoutes(a.srv, a.db)
}

func (a *Api) Run(bindAddress string) error {
	return a.srv.Run(bindAddress)
}
