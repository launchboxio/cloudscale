package listener

import (
	"github.com/launchboxio/cloudscale/internal/proxy/backend"
	"net/http"
)

type HttpListenerOptions struct {
}

type HttpListener struct {
	baseListener
}

func NewHttpListener(opts *HttpListenerOptions) *HttpListener {
	return &HttpListener{}
}

func (t *HttpListener) Send(dst *backend.Backend) {
	peer := lb.serverPool.GetNextValidPeer()
	if peer != nil {
		peer.Serve(w, r)
		return
	}
	http.Error(w, "Service not available", http.StatusServiceUnavailable)
}
