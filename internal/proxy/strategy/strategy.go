package strategy

import "github.com/launchboxio/cloudscale/internal/proxy/backend"

type baseStrategy struct {
	Backends []*backend.Backend
}

type Strategy interface {
	Decide() (*backend.Backend, error)
}
