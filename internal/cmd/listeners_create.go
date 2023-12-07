package cmd

import (
	"github.com/launchboxio/cloudscale/sdk"
	"github.com/spf13/cobra"
	"log"
)

var listenersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new listener",
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()

		name, _ := flags.GetString("name")
		ipAddr, _ := flags.GetString("ip-addr")
		port, _ := flags.GetUint16("port")
		lbType, _ := flags.GetString("type")

		input := &sdk.CreateListenerInput{
			Name:      name,
			IpAddress: ipAddr,
			Port:      port,
			Type:      lbType,
		}

		res, err := client.Listeners().Create(input)
		if err != nil {
			log.Fatal(err)
		}
		print(res.Listener)
	},
}

func init() {
	listenersCreateCmd.Flags().String("name", "", "Name of the listener")
	_ = listenersCreateCmd.MarkFlagRequired("name")

	listenersCreateCmd.Flags().String("ip-address", "", "IP address to bind to")
	_ = listenersCreateCmd.MarkFlagRequired("ip-address")

	listenersCreateCmd.Flags().Uint16("port", 0, "Port to bind to")
	_ = listenersCreateCmd.MarkFlagRequired("port")

	listenersCreateCmd.Flags().String("type", "", "Type of load balancer")
	_ = listenersCreateCmd.MarkFlagRequired("type")

	listenersCmd.AddCommand(listenersCreateCmd)
}
