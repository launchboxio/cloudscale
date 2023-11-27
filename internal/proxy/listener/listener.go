package listener

import (
	"crypto/tls"
	"github.com/launchboxio/cloudscale/internal/proxy/backend"
	"github.com/launchboxio/cloudscale/internal/proxy/targetgroup"
)

type Listener interface {
	Send(b *backend.Backend)
}

type baseListener struct {
	Port     uint16
	Protocol string
	Tls      tls.Config

	Default *targetgroup.TargetGroup
}
