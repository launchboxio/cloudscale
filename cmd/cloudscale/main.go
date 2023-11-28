package main

import (
	"github.com/launchboxio/cloudscale/internal/controller"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "cloudscale",
	Short: "Cloudscale Load Balancer",
	Run: func(cmd *cobra.Command, args []string) {
		ctrl := controller.Controller{}

		if err := ctrl.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
