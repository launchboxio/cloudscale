package listener

import (
	"github.com/launchboxio/cloudscale/internal/proxy/backend"
)

type TcpListenerOptions struct {
}

type TcpListener struct {
	baseListener
}

func NewTcpListener(opts *TcpListenerOptions) *TcpListener {
	return &TcpListener{}
}

func (t *TcpListener) Send(dst *backend.Backend) {

}
