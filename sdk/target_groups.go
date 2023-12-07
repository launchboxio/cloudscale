package sdk

import "github.com/launchboxio/cloudscale/internal/api"

type TargetGroups struct {
	*Client
}

type TargetGroupsList struct {
	TargetGroups []api.TargetGroup `json:"target_groups"`
}

type TargetGroupResponse struct {
	TargetGroup api.TargetGroup `json:"target_group"`
}

func (t *TargetGroups) List() (TargetGroupsList, error) {
	var targetGroupList TargetGroupsList
	_, err := t.http.R().
		SetResult(&targetGroupList).
		Get("/target_groups")
	return targetGroupList, err
}

type TargetGroupAttachmentInput struct {
	TargetGroupId string `json:"-"`
	IpAddress     string `json:"ip_address"`
	Port          uint16 `json:"port"`
}

func (t *TargetGroups) AddAttachment(input TargetGroupAttachmentInput) (TargetGroupResponse, error) {
	var response TargetGroupResponse
	_, err := t.http.R().
		SetResult(&response).
		Post("/target_groups/" + input.TargetGroupId + "/attachments")
	return response, err
}
