package backend

import (
	"net"
	"sync"
)

type Backend struct {
	IpAddress net.IPAddr
	Port      uint16

	// connections counter
	connections int
	mux         sync.RWMutex

	// Whether backend instance is healthy
	healthy bool

	// Whether backend instance is ready
	ready bool
}
