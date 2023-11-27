package strategy

import (
	"github.com/launchboxio/cloudscale/internal/proxy/backend"
	"sync"
)

type RoundRobinOptions struct{}

type RoundRobin struct {
	baseStrategy

	mux            sync.RWMutex
	currentBackend *backend.Backend
}

func (rr *RoundRobin) Decide() (*backend.Backend, error) {
	return rr.Backends[0], nil
}
