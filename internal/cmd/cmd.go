package cmd

import (
  "encoding/json"
  "fmt"
  "github.com/launchboxio/cloudscale/sdk"
  "github.com/spf13/cobra"
  "log"
)

var client *sdk.Client

var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "Manage Cloudscale via API",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		endpoint, _ := cmd.Flags().GetString("cloudscale-addr")
		client = sdk.New(endpoint)
	},
}

func init() {
	ApiCmd.PersistentFlags().String("cloudscale-addr", "http://localhost:9001", "Cloudscale HTTP Endpoint")
}

func print(data interface{}) {
	out, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
