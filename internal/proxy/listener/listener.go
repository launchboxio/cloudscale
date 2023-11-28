package listener

import (
	"crypto/tls"
	"github.com/launchboxio/cloudscale/internal/proxy/backend"
	"github.com/launchboxio/cloudscale/internal/proxy/features/stickiness"
	"github.com/launchboxio/cloudscale/internal/proxy/targetgroup"
	"sync"
)

type ActionType string

const (
	ActionTypeForward  ActionType = "forward"
	ActionTypeRedirect            = "redirect"
)

type Listener interface {
	Send(b *backend.Backend)
}

type baseListener struct {
	Port     uint16
	Protocol string
	Tls      tls.Config

	Default *targetgroup.TargetGroup
	Rules   Rule

	mux    sync.Mutex
	closed bool
	cond   *sync.Cond
}

type Rule struct {
	Priority uint16
	Action   Action
}

type Condition struct {
	HostHeader        []string
	HttpHeader        []string
	HttpRequestMethod []string
	PathPattern       []string
	SourceIp          []string
}

type Action struct {
	Type     string
	Forward  ForwardAction
	Redirect RedirectAction
}

type ForwardAction struct {
	TargetGroup TargetGroupForwardAction
	Stickiness  stickiness.Stickiness
}

type TargetGroupForwardAction struct {
	TargetGroup targetgroup.TargetGroup
	Weight      uint8
}

type RedirectAction struct {
	Host       string
	Port       string
	Path       string
	Protocol   string
	Query      string
	StatusCode string
}
