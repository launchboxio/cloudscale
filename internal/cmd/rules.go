package cmd

import "github.com/spf13/cobra"

var rulesCmd = &cobra.Command{
	Use:   "rules",
	Short: "Interact with CloudScale listener rules",
}

func init() {
	ApiCmd.AddCommand(rulesCmd)
}
