package main

import (
  "github.com/spf13/cobra"
  "log"
)

var rootCmd = &cobra.Command{
	Use:   "cloudscale",
	Short: "Cloudscale Load Balancer",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
