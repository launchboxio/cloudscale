package strategy

import (
	"github.com/launchboxio/cloudscale/internal/proxy/backend"
	"sync"
)

type RoundRobinOptions struct{}

type RoundRobin struct {
	baseStrategy

	mux            sync.RWMutex
	currentBackend int
}

func (rr *RoundRobin) Decide() (*backend.Backend, error) {
	for i := 0; i < len(rr.Backends); i++ {
		next := rr.selectNext()
		if next.IsHealthy() {
			return next, nil
		}
	}
	return nil, nil
}

func (rr *RoundRobin) selectNext() *backend.Backend {
	rr.withLock(func() {
		rr.currentBackend = (rr.currentBackend + 1) % len(rr.Backends)
	})
	return rr.Backends[rr.currentBackend]
}

func (rr *RoundRobin) withLock(f func()) {
	rr.mux.Lock()
	f()
	rr.mux.Unlock()
}
