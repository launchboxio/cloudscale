package cmd

import "github.com/spf13/cobra"

var attachmentsCmd = &cobra.Command{
	Use:   "attachments",
	Short: "Interact with CloudScale target group attachments",
}

func init() {
	ApiCmd.AddCommand(attachmentsCmd)
}
