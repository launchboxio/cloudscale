package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var certificatesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available SSL certificates",
	Run: func(cmd *cobra.Command, args []string) {
		certificates, err := client.Certificates().List()
		if err != nil {
			log.Fatal(err)
		}
		print(certificates.Certificates)
	},
}

func init() {
	certificatesCmd.AddCommand(certificatesListCmd)
}
