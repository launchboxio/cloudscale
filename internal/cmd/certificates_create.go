package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var certificatesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Import a new TLS certificates",
	Run: func(cmd *cobra.Command, args []string) {
		certPath, _ := cmd.Flags().GetString("cert")
		keyPath, _ := cmd.Flags().GetString("key")
		if certPath == "" || keyPath == "" {
			log.Fatal("Both --key and --cert flags are required")
		}

		certData, err := os.ReadFile(certPath)
		if err != nil {
			log.Fatal("Failed reading certificate file")
		}

		key, err := os.ReadFile(keyPath)
		if err != nil {
			log.Fatal(err)
		}

		// Post the data to create the certificate
		fmt.Println(key)
		fmt.Println(certData)
	},
}

func init() {
	certificatesCreateCmd.Flags().String("cert", "", "Path to the TLS certificate")
	certificatesCreateCmd.Flags().String("key", "", "Path to the TLS private key")

	_ = certificatesCreateCmd.MarkFlagRequired("cert")
	_ = certificatesCreateCmd.MarkFlagRequired("key")

	certificatesCmd.AddCommand(certificatesCreateCmd)
}
