package listener

import "github.com/launchboxio/cloudscale/internal/proxy/backend"

type HttpListenerOptions struct {
}

type HttpListener struct {
	baseListener
}

func NewHttpListener(opts *HttpListenerOptions) *HttpListener {
	return &HttpListener{}
}

func (t *HttpListener) Send(dst *backend.Backend) {

}
