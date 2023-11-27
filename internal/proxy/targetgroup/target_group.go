package targetgroup

import (
	"github.com/launchboxio/cloudscale/internal/proxy/backend"
	"github.com/launchboxio/cloudscale/internal/proxy/features/stickiness"
	"github.com/launchboxio/cloudscale/internal/proxy/strategy"
)

type TargetGroup struct {
	// Port to direct traffic on
	Port uint16

	// Protocol of the target group
	Protocol string

	// Strategy for load balancing decisions
	Strategy strategy.Strategy

	// Stickiness configures the stickiness feature
	Stickiness stickiness.Stickiness

	// Backends contain all registered members for this target group
	Backends []*backend.Backend
}

func (tg *TargetGroup) GetNextAvailable() (*backend.Backend, error) {
	return tg.Strategy.Decide()
}
