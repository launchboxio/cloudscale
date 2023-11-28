package main

import (
	"github.com/launchboxio/cloudscale/internal/controller"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "Start the control plane for CloudScale",
	Run: func(cmd *cobra.Command, args []string) {
		grpcBindAddress, _ := cmd.Flags().GetString("grpc-bind-address")
		httpBindAddress, _ := cmd.Flags().GetString("http-bind-address")

		ctrl := controller.New(&controller.Options{
			GrpcAddress: grpcBindAddress,
			HttpAddress: httpBindAddress,
		})

		if err := ctrl.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	controllerCmd.Flags().String("grpc-bind-address", ":10001", "Address to bind the Envoy GRPC Controller to")
	controllerCmd.Flags().String("http-bind-address", ":9001", "Address to bind the HTTP API server to")
	rootCmd.AddCommand(controllerCmd)
}
