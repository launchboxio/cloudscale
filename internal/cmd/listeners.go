package cmd

import "github.com/spf13/cobra"

var listenersCmd = &cobra.Command{
	Use:   "listeners",
	Short: "Interact with CloudScale listeners",
}

func init() {
	ApiCmd.AddCommand(listenersCmd)
}
