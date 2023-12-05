package cmd

import "github.com/spf13/cobra"

var certificatesCmd = &cobra.Command{
	Use:   "certificates",
	Short: "Interact with CloudScale certificates",
}

func init() {
	ApiCmd.AddCommand(certificatesCmd)
}
