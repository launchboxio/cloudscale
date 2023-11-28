package strategy

import (
	"github.com/launchboxio/cloudscale/internal/proxy/backend"
	"sort"
)

type LeastConnections struct {
	baseStrategy
}

func (lc *LeastConnections) Decide() (*backend.Backend, error) {
	sort.Slice(lc.Backends, func(i, j int) bool {
		return lc.Backends[i].CurrentConnections() < lc.Backends[j].CurrentConnections()
	})
	return lc.Backends[0], nil
}
