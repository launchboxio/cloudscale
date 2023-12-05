package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var listenersGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve an individual listener",
	Run: func(cmd *cobra.Command, args []string) {
		listenerId, _ := cmd.Flags().GetString("id")
		listener, err := client.Listeners().Get(listenerId)
		if err != nil {
			log.Fatal(err)
		}

		print(listener.Listener)
	},
}

func init() {
	listenersGetCmd.Flags().String("id", "", "ID of the listener")

	_ = listenersGetCmd.MarkFlagRequired("id")

	listenersCmd.AddCommand(listenersGetCmd)
}
