package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var listListenersCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured listeners",
	Run: func(cmd *cobra.Command, args []string) {
		listeners, err := client.Listeners().List()
		if err != nil {
			log.Fatal(err)
		}

		print(listeners.Listeners)
	},
}

func init() {
	listenersCmd.AddCommand(listListenersCmd)
}
