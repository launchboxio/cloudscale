package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
	"net/http"
)

const CertificateBucket = "Certificates"

type certificateCtrl struct {
	db *bolt.DB
}

type Certificate struct {
	Id   string `json:"id"`
	Cert []byte `json:"cert"`
	Key  []byte `json:"key"`
}

func registerCertificateRoutes(r *gin.Engine, db *bolt.DB) {
	ctrl := certificateCtrl{db}
	c := r.Group("/certificates")
	c.GET("", ctrl.list)
	c.POST("", ctrl.create)
	c.GET("/:certificateId", ctrl.show)
	c.PUT("/:certificateId", ctrl.update)
	c.DELETE("/:certificateId", ctrl.delete)
}

func (ctrl *certificateCtrl) list(c *gin.Context) {
	var certificates = []Certificate{}
	err := ctrl.db.View(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(CertificateBucket)).ForEach(func(k, v []byte) error {
			var certificate Certificate
			if err := json.Unmarshal(v, &certificate); err != nil {
				return err
			}
			certificate.Id = string(k)
			certificates = append(certificates, certificate)
			return nil
		})
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"certificates": certificates})
}

func (ctrl *certificateCtrl) show(c *gin.Context) {
	var certificate *Certificate

	err := ctrl.db.View(func(tx *bolt.Tx) error {
		data := tx.Bucket([]byte(CertificateBucket)).Get([]byte(c.Param("certificateId")))
		if data == nil {
			return nil
		}

		certificate = &Certificate{}
		err := json.Unmarshal(data, certificate)
		if err != nil {
			return err
		}

		certificate.Id = c.Param("certificateId")
		return nil
	})

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
	var input Certificate
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
		return tx.Bucket([]byte(CertificateBucket)).Put([]byte(input.Id), data)
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"certificate": input})
}

func (ctrl *certificateCtrl) update(c *gin.Context) {
	certificateId := c.Param("certificateId")
	var input Certificate

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Id = certificateId
	err := ctrl.db.Update(func(tx *bolt.Tx) error {
		data, err := json.Marshal(input)
		if err != nil {
			return err
		}

		return tx.Bucket([]byte(CertificateBucket)).Put([]byte(certificateId), data)
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"certificate": input})
}

func (ctrl *certificateCtrl) delete(c *gin.Context) {
	certificateId := c.Param("certificateId")
	err := ctrl.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(CertificateBucket)).Delete([]byte(certificateId))
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
