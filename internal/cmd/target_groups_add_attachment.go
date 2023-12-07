package cmd

import (
	"github.com/launchboxio/cloudscale/sdk"
	"github.com/spf13/cobra"
	"log"
)

var targetGroupsAddAttachmentCmd = &cobra.Command{
	Use:   "add-attachment",
	Short: "Add an upstream node to a target group",
	Run: func(cmd *cobra.Command, args []string) {
		targetGroupId, _ := cmd.Flags().GetString("target-group-id")
		ipAddr, _ := cmd.Flags().GetString("ip-address")
		port, _ := cmd.Flags().GetUint16("port")

		input := sdk.TargetGroupAttachmentInput{
			TargetGroupId: targetGroupId,
			IpAddress:     ipAddr,
			Port:          port,
		}

		res, err := client.TargetGroups().AddAttachment(input)
		if err != nil {
			log.Fatal(err)
		}

		print(res.TargetGroup)
	},
}

func init() {
	targetGroupsAddAttachmentCmd.Flags().String("target-group-id", "", "ID of the target group")
	_ = targetGroupsAddAttachmentCmd.MarkFlagRequired("target-group-id")

	targetGroupsAddAttachmentCmd.Flags().String("ip-address", "", "IP Address of the upstream")
	_ = targetGroupsAddAttachmentCmd.MarkFlagRequired("ip-address")

	targetGroupsAddAttachmentCmd.Flags().Uint16("port", 0, "Port of the upstream")
	_ = targetGroupsAddAttachmentCmd.MarkFlagRequired("port")

	targetGroupsCmd.AddCommand(targetGroupsAddAttachmentCmd)
}
