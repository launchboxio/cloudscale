package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
	"net/http"
)

const ListenersBucket = "Listeners"

type listenerCtrl struct {
	db *bolt.DB
}

type Listener struct {
	Id       string `json:"id,omitempty"`
	Port     uint16 `json:"port"`
	Protocol string `json:"protocol,omitempty"`
	//Tls      tls.Config `json:"tls,omitempty"`

	Rules []Rule `json:"rules"`
}

type Action struct {
	Type string `json:"type"`
	//Forward  ForwardAction  `json:"forward,omitempty"`
	//Redirect RedirectAction `json:"redirect,omitempty"`
}

//
//type ForwardAction struct {
//	TargetGroup TargetGroupForwardAction `json:"target_group"`
//	Stickiness  stickiness.Stickiness    `json:"stickiness,omitempty"`
//}
//
//type TargetGroupForwardAction struct {
//	TargetGroup targetgroup.TargetGroup `json:"target_group"`
//	Weight      uint8                   `json:"weight,omitempty"`
//}

type RedirectAction struct {
	Host       string `json:"host"`
	Port       string `json:"port,omitempty"`
	Path       string `json:"path,omitempty"`
	Protocol   string `json:"protocol,omitempty"`
	Query      string `json:"query,omitempty"`
	StatusCode string `json:"status_code,omitempty"`
}

type Rule struct {
	Priority uint16 `json:"priority"`
	Action   Action `json:"action"`
}

type Condition struct {
	HostHeader        []string `json:"host_header,omitempty"`
	HttpHeader        []string `json:"http_header,omitempty"`
	HttpRequestMethod []string `json:"http_request_method,omitempty"`
	PathPattern       []string `json:"path_pattern,omitempty"`
	SourceIp          []string `json:"source_ip,omitempty"`
}

func registerListenerRoutes(r *gin.Engine, db *bolt.DB) {
	ctrl := listenerCtrl{db}
	l := r.Group("/listeners")
	l.GET("", ctrl.list)
	l.POST("", ctrl.create)
	l.GET("/:listenerId", ctrl.show)
	l.PUT("/:listenerId", ctrl.update)
	l.DELETE("/:listenerId", ctrl.delete)
}

func (ctrl *listenerCtrl) list(c *gin.Context) {
	var listeners []Listener
	if ctrl.db == nil {
		fmt.Println("Database not initialized")
	}
	err := ctrl.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(ListenersBucket))
		return b.ForEach(func(k, v []byte) error {
			var listener Listener
			if err := json.Unmarshal(v, &listener); err != nil {
				return err
			}
			listener.Id = string(k)
			listeners = append(listeners, listener)
			return nil
		})
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if listeners == nil {
		c.JSON(http.StatusOK, gin.H{"listeners": []Listener{}})
	} else {
		c.JSON(http.StatusOK, gin.H{"listeners": listeners})
	}
}

func (ctrl *listenerCtrl) show(c *gin.Context) {
	var listener *Listener

	err := ctrl.db.View(func(tx *bolt.Tx) error {
		data := tx.Bucket([]byte(ListenersBucket)).Get([]byte(c.Param("listenerId")))
		if data == nil {
			return nil
		}
		listener = &Listener{}
		err := json.Unmarshal(data, listener)
		if err != nil {
			return err
		}
		listener.Id = c.Param("listenerId")
		return nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if listener == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"listener": listener})
}

func (ctrl *listenerCtrl) create(c *gin.Context) {
	var input Listener
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
		return tx.Bucket([]byte(ListenersBucket)).Put([]byte(input.Id), data)
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"listener": input})
}

func (ctrl *listenerCtrl) update(c *gin.Context) {
	listenerId := c.Param("listenerId")
	var input Listener

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Id = listenerId
	err := ctrl.db.Update(func(tx *bolt.Tx) error {
		data, err := json.Marshal(input)
		if err != nil {
			return err
		}

		return tx.Bucket([]byte(ListenersBucket)).Put([]byte(listenerId), data)
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"listener": input})
}

func (ctrl *listenerCtrl) delete(c *gin.Context) {
	listenerId := c.Param("listenerId")
	err := ctrl.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(ListenersBucket)).Delete([]byte(listenerId))
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
