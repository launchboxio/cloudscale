package api

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

type targetGroupCtrl struct {
	db *Service
}

type TargetGroup struct {
	Base
	Name        string                  `json:"name"`
	Enabled     bool                    `json:"enabled,omitempty"`
	Attachments []TargetGroupAttachment `json:"attachments"`
}

type TargetGroupAttachment struct {
	Base
	Id            string `json:"id"`
	IpAddress     net.IP `json:"ip_address"`
	Port          uint16 `json:"port"`
	TargetGroupID string
}

func registerTargetGroupRoutes(r *gin.Engine, db *Service) {
	ctrl := targetGroupCtrl{db}
	c := r.Group("/target_groups")
	c.GET("", ctrl.list)
	c.POST("", ctrl.create)
	c.GET("/:targetGroupId", ctrl.show)
	c.PUT("/:targetGroupId", ctrl.update)
	c.DELETE("/:targetGroupId", ctrl.delete)
}

func (ctrl *targetGroupCtrl) list(c *gin.Context) {
	targetGroups, err := ctrl.db.ListTargetGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"target_groups": targetGroups})
}

func (ctrl *targetGroupCtrl) show(c *gin.Context) {
	targetGroup, err := ctrl.db.GetTargetGroup(c.Param("targetGroupId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if targetGroup == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"target_group": targetGroup})
}

func (ctrl *targetGroupCtrl) create(c *gin.Context) {
	var input *TargetGroup
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group, err := ctrl.db.CreateTargetGroup(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"target_group": group})
}

func (ctrl *targetGroupCtrl) update(c *gin.Context) {
	targetGroupId := c.Param("targetGroupId")
	var input *TargetGroup

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := ctrl.db.GetTargetGroup(targetGroupId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	group, err := ctrl.db.UpdateTargetGroup(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"target_group": group})
}

func (ctrl *targetGroupCtrl) delete(c *gin.Context) {
	targetGroupId := c.Param("targetGroupId")
	if err := ctrl.db.DestroyTargetGroup(targetGroupId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
