package sdk

import (
	"fmt"
	"github.com/launchboxio/cloudscale/internal/api"
)

type Certificates struct {
	*Client
}

type CertificatesList struct {
	Certificates []api.Certificate `json:"certificates"`
}

type CertificateResponse struct {
	Certificate *api.Certificate `json:"certificate"`
}

func (c *Certificates) List() (CertificatesList, error) {
	var certList CertificatesList
	_, err := c.http.R().
		SetResult(&certList).
		Get("/certificates")
	return certList, err
}

type CreateCertificateInput struct {
	Certificate []byte `json:"cert"`
	Key         []byte `json:"key"`
	Name        string `json:"name"`
}

func (c *Certificates) Create(input *CreateCertificateInput) (*CertificateResponse, error) {
	var res *CertificateResponse
	resp, err := c.http.R().
		SetBody(input).
		SetResult(&res).
		Post("/certificates")
	fmt.Println(resp)
	return res, err
}
