package cmd

import "github.com/spf13/cobra"

var targetGroupsCmd = &cobra.Command{
	Use:   "target-groups",
	Short: "Interact with CloudScale target groups / clusters",
}

func init() {
	ApiCmd.AddCommand(targetGroupsCmd)
}
