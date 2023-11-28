package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"github.com/google/uuid"
	"github.com/launchboxio/cloudscale/sdk"
	"log"
)

func main() {
	client := sdk.New("http://localhost:9001")

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}
	certs, err := client.Certificates().Create(&sdk.CreateCertificateInput{
		Certificate: priv.N.Bytes(),
		Key:         priv.D.Bytes(),
		Name:        uuid.New().String(),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(certs)
}
