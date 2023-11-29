package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type certificateCtrl struct {
	db *Service
}

type Certificate struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name"`
	Cert []byte `json:"cert"`
	Key  []byte `json:"key"`
}

func registerCertificateRoutes(r *gin.Engine, db *Service) {
	ctrl := certificateCtrl{db}
	c := r.Group("/certificates")
	c.GET("", ctrl.list)
	c.POST("", ctrl.create)
	c.GET("/:certificateId", ctrl.show)
	c.PUT("/:certificateId", ctrl.update)
	c.DELETE("/:certificateId", ctrl.delete)
}

func (ctrl *certificateCtrl) list(c *gin.Context) {
	certificates, err := ctrl.db.ListCertificates()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"certificates": certificates})
}

func (ctrl *certificateCtrl) show(c *gin.Context) {
	certificate, err := ctrl.db.GetCertificate(c.Param("certificateId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if certificate == nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"certificate": certificate})
}

func (ctrl *certificateCtrl) create(c *gin.Context) {
	var input *Certificate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	certificate, err := ctrl.db.CreateCertificate(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"certificate": certificate})
}

func (ctrl *certificateCtrl) update(c *gin.Context) {
	certificateId := c.Param("certificateId")
	var input *Certificate

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Id = certificateId
	certificate, err := ctrl.db.UpdateCertificate(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"certificate": certificate})
}

func (ctrl *certificateCtrl) delete(c *gin.Context) {
	certificateId := c.Param("certificateId")

	if err := ctrl.db.DeleteCertificate(certificateId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
