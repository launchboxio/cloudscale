package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
	"net/http"
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
	c.DELETE("/:targetGroupId", ctrl.delete)
}

func (ctrl *targetGroupCtrl) list(c *gin.Context) {
	var targetGroups = []TargetGroup{}
	err := ctrl.db.View(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(TargetGroupsBucket)).ForEach(func(k, v []byte) error {
			var targetGroup TargetGroup
			if err := json.Unmarshal(v, &targetGroup); err != nil {
				return err
			}
			targetGroup.Id = string(k)
			targetGroups = append(targetGroups, targetGroup)
			return nil
		})
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"target_groups": targetGroups})
}

func (ctrl *targetGroupCtrl) show(c *gin.Context) {
	var targetGroup *TargetGroup
	err := ctrl.db.View(func(tx *bolt.Tx) error {
		data := tx.Bucket([]byte(TargetGroupsBucket)).Get([]byte(c.Param("targetGroupId")))
		if data == nil {
			return nil
		}

		targetGroup = &TargetGroup{}
		err := json.Unmarshal(data, targetGroup)
		if err != nil {
			return err
		}

		targetGroup.Id = c.Param("targetGroupId")
		return nil
	})

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
	var input TargetGroup
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Id = uuid.New().String()
	err := ctrl.db.Update(func(tx *bolt.Tx) error {
		data, err := json.Marshal(input)
		if err != nil {
			return err
		}

		return tx.Bucket([]byte(TargetGroupsBucket)).Put([]byte(input.Id), data)
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"target_group": input})
}

func (ctrl *targetGroupCtrl) update(c *gin.Context) {
	targetGroupId := c.Param("targetGroupId")
	var input TargetGroup

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Id = targetGroupId
	err := ctrl.db.Update(func(tx *bolt.Tx) error {
		data, err := json.Marshal(input)
		if err != nil {
			return err
		}

		return tx.Bucket([]byte(TargetGroupsBucket)).Put([]byte(targetGroupId), data)
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"target_group": input})
}

func (ctrl *targetGroupCtrl) delete(c *gin.Context) {
	targetGroupId := c.Param("targetGroupId")
	err := ctrl.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(TargetGroupsBucket)).Delete([]byte(targetGroupId))
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
