package cmd

import "github.com/spf13/cobra"

var addListenerRuleCmd = &cobra.Command{
	Use:   "add-rule",
	Short: "Add a new rule for a listener",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	addListenerRuleCmd.Flags().String("listener-id", "", "ID of the listener")
	_ = addListenerRuleCmd.MarkFlagRequired("listener-id")

	addListenerRuleCmd.Flags().Uint16("priority", 0, "Priority of the listener rule")
	addListenerRuleCmd.Flags().String("action", "", "Action for this rule")

	listenersCmd.AddCommand(addListenerRuleCmd)
}
