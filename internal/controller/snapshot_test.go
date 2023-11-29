package controller

import (
	"fmt"
	"github.com/launchboxio/cloudscale/internal/api"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
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
			},
		},
		TargetGroupAttachments: []*api.TargetGroupAttachment{
			{
				Id:            "zzzzzzz-zz-zz-zzzzz",
				TargetGroupId: "yyyyyyy-yy-yy-yyyyy",
				IpAddress:     net.ParseIP("172.10.0.1"),
				Port:          3000,
			},
		},
	}
	snapshot, _ := generateSnapshot(info)
	fmt.Println(snapshot)

	out, err := yaml.Marshal(snapshot.VersionMap)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}
