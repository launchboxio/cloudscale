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

func (tl *TcpListener) Close() error {
	tl.WithLock(func() {
		// TODO: Check if listener closed
		if tl.closed {
			return nil
		}
	})
	tl.closed = true
	tl.cond.Broadcast()
	return nil

	tl.lock()
	if tl.closed {
		tl.mu.Unlock()
		return nil
	}
	tl.closed = true
	tl.mu.Unlock()
	tl.cond.Broadcast()
	return nil
}

func (tl *TcpListener) WithLock(f func()) {
	tl.mux.Lock()
	f()
	tl.mux.Unlock()
}
