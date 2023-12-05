package cmd

import "github.com/spf13/cobra"

var listenersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new listener",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	listenersCreateCmd.Flags().String("name", "", "Name of the listener")
	listenersCreateCmd.Flags().String("ip_address", "", "IP address to bind to")
	listenersCreateCmd.Flags().Uint16("port", 0, "Port to bind to")
	listenersCreateCmd.Flags().String("type", "", "Type of load balancer")

	listenersCmd.AddCommand(listenersCreateCmd)
}
