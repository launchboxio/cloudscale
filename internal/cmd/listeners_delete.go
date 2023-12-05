package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var listenersDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a listener",
	Run: func(cmd *cobra.Command, args []string) {
		listenerId, _ := cmd.Flags().GetString("id")
		if err := client.Listeners().Delete(listenerId); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Listener deleted")
	},
}

func init() {
	listenersDeleteCmd.Flags().String("id", "", "ID of the listener")
	_ = listenersDeleteCmd.MarkFlagRequired("id")

	listenersCmd.AddCommand(listenersDeleteCmd)
}
