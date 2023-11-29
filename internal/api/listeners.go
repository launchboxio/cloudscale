package api

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

type listenerCtrl struct {
	db *Service
}

type Listener struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name"`
	IpAddress net.IP `json:"ip_address,omitempty"`
	Port      uint16 `json:"port"`
	Protocol  string `json:"protocol,omitempty"`

	SslCertificate string `json:"ssl_certificate,omitempty"`

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

func registerListenerRoutes(r *gin.Engine, db *Service) {
	ctrl := listenerCtrl{db}
	l := r.Group("/listeners")
	l.GET("", ctrl.list)
	l.POST("", ctrl.create)
	l.GET("/:listenerId", ctrl.show)
	l.PUT("/:listenerId", ctrl.update)
	l.DELETE("/:listenerId", ctrl.delete)
}

func (ctrl *listenerCtrl) list(c *gin.Context) {
	listeners, err := ctrl.db.ListListeners()

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
	listener, err := ctrl.db.GetListener(c.Param("listenerId"))
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
	var input *Listener
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	listener, err := ctrl.db.CreateListener(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"listener": listener})
}

func (ctrl *listenerCtrl) update(c *gin.Context) {
	listenerId := c.Param("listenerId")
	var input *Listener

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Id = listenerId
	listener, err := ctrl.db.UpdateListener(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"listener": listener})
}

func (ctrl *listenerCtrl) delete(c *gin.Context) {
	listenerId := c.Param("listenerId")

	if err := ctrl.db.DestroyListener(listenerId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
