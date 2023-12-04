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
				Type:      "layer7",
				IpAddress: net.ParseIP("127.0.0.1"),
				Rules: []api.Rule{
					{
						Action: api.Action{
							Type: "forward",
							Forward: api.ForwardAction{
								TargetGroup: api.TargetGroupForwardAction{
									TargetGroupId: "yyyyyyy-yy-yy-yyyyy",
								},
							},
						},
					},
				},
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
	//fmt.Println(snapshot)
	fmt.Println(snapshot.Resources[0])
	assert.Nil(t, err)
	assert.NoError(t, snapshot.Consistent())
}

func Test_Layer4TcpListener(t *testing.T) {
	info := &SnapshotInfo{
		Listeners: []*api.Listener{{
			Id:        "xxxxxx-xx-xx-xxxxx",
			Name:      "test-listener",
			Type:      "layer4",
			IpAddress: net.ParseIP("127.0.0.1"),
			Rules: []api.Rule{
				{
					Action: api.Action{
						Type: "forward",
						Forward: api.ForwardAction{
							TargetGroup: api.TargetGroupForwardAction{
								TargetGroupId: "yyyyyyy-yy-yy-yyyyy",
							},
						},
					},
				},
			},
		}},
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
	assert.Nil(t, err)
	assert.NoError(t, snapshot.Consistent())
}
