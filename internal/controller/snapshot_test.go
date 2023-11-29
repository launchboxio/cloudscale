package controller

import (
	"fmt"
	"github.com/launchboxio/cloudscale/internal/api"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func Test_Snapshot(t *testing.T) {
	info := &SnapshotInfo{
		Listeners: []*api.Listener{
			{
				Id:        "xxxxxx-xx-xx-xxxxx",
				Name:      "test-listener",
				IpAddress: net.ParseIP("127.0.0.1"),
			},
		},
		TargetGroups: []*api.TargetGroup{
			{
				Id:   "yyyyyyy-yy-yy-yyyyy",
				Name: "test-target-group",
				Attachments: []api.TargetGroupAttachment{
					{
						Id:        "zzzzzzz-zz-zz-zzzzz",
						IpAddress: net.ParseIP("172.10.0.1"),
						Port:      3000,
					},
				},
			},
		},
	}
	snapshot, err := generateSnapshot(info)
	fmt.Println(snapshot)
	assert.Nil(t, err)
	assert.NoError(t, snapshot.Consistent())
}
