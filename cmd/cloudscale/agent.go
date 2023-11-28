package main

import (
	"github.com/spf13/cobra"
	"os"
)

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Agent for sending resource metrics to controller",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Agent functionality not yet implemented")
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)
}
