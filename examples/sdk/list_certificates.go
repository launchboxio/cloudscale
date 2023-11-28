package main

import (
	"fmt"
	"github.com/launchboxio/cloudscale/sdk"
	"log"
)

func main() {
	client := sdk.New("http://localhost:9001")

	certs, err := client.Certificates().List()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(certs.Certificates)
}
