package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var targetGroupsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured target groups",
	Run: func(cmd *cobra.Command, args []string) {
		targetGroups, err := client.TargetGroups().List()
		if err != nil {
			log.Fatal(err)
		}
		print(targetGroups.TargetGroups)
	},
}

func init() {
	targetGroupsCmd.AddCommand(targetGroupsListCmd)
}
