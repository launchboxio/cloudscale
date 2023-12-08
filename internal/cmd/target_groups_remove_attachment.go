package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var targetGroupsRemoveAttachmentCmd = &cobra.Command{
	Use:   "remove-attachment",
	Short: "Remove an upstream from the target group",
	Run: func(cmd *cobra.Command, args []string) {
		targetGroupId, _ := cmd.Flags().GetString("target-group-id")
		attachmentId, _ := cmd.Flags().GetString("attachment-id")

		if err := client.TargetGroups().RemoveAttachment(targetGroupId, attachmentId); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Attachment removed from target group")
	},
}

func init() {
	targetGroupsRemoveAttachmentCmd.Flags().String("target-group-id", "", "ID of the target group")
	_ = targetGroupsRemoveAttachmentCmd.MarkFlagRequired("target-group-id")

	targetGroupsRemoveAttachmentCmd.Flags().String("attachment-id", "", "ID of the attachment")
	_ = targetGroupsRemoveAttachmentCmd.MarkFlagRequired("attachment-id")

	targetGroupsCmd.AddCommand(targetGroupsRemoveAttachmentCmd)
}
