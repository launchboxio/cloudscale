package api

import (
	"github.com/gin-gonic/gin"
	bolt "go.etcd.io/bbolt"
)

const TargetGroupsBucket = "TargetGroups"

type targetGroupCtrl struct {
	db *bolt.DB
}

type TargetGroup struct {
	Id string `json:"id"`
}

func registerTargetGroupRoutes(r *gin.Engine, db *bolt.DB) {
	ctrl := targetGroupCtrl{db}
	c := r.Group("/target_groups")
	c.GET("", ctrl.list)
	c.POST("", ctrl.create)
	c.GET("/:targetGroupId", ctrl.show)
	c.PUT("/:targetGroupId", ctrl.update)
	c.DELETE("/:certificateId", ctrl.delete)
}

func (ctrl *targetGroupCtrl) list(c *gin.Context) {

}

func (ctrl *targetGroupCtrl) show(c *gin.Context) {

}

func (ctrl *targetGroupCtrl) create(c *gin.Context) {

}

func (ctrl *targetGroupCtrl) update(c *gin.Context) {

}

func (ctrl *targetGroupCtrl) delete(c *gin.Context) {

}
