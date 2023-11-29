package controller

import (
	"context"
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discoverygrpc "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	log "github.com/sirupsen/logrus"
	"sync"
)

func (cb *Callbacks) Report() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	log.WithFields(log.Fields{"fetches": cb.Fetches, "requests": cb.Requests}).Info("cb.Report()  callbacks")
}
func (cb *Callbacks) OnStreamOpen(_ context.Context, id int64, typ string) error {
	log.Infof("OnStreamOpen %d open for %s", id, typ)
	return nil
}
func (cb *Callbacks) OnStreamClosed(id int64, node *corev3.Node) {
	log.Infof("OnStreamClosed %d closed", id)
}

func (cb *Callbacks) OnStreamRequest(id int64, r *discoverygrpc.DiscoveryRequest) error {
	log.Infof("OnStreamRequest %v", r.TypeUrl)
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.Requests++
	if cb.Signal != nil {
		close(cb.Signal)
		cb.Signal = nil
	}
	return nil
}
func (cb *Callbacks) OnStreamResponse(ctx context.Context, id int64, req *discoverygrpc.DiscoveryRequest, resp *discoverygrpc.DiscoveryResponse) {
	log.Infof("OnStreamResponse... %d   Request [%v],  Response[%v]", id, req.TypeUrl, resp.TypeUrl)
	cb.Report()
}

func (cb *Callbacks) OnFetchRequest(ctx context.Context, req *discoverygrpc.DiscoveryRequest) error {
	log.Infof("OnFetchRequest... Request [%v]", req.TypeUrl)
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.Fetches++
	if cb.Signal != nil {
		close(cb.Signal)
		cb.Signal = nil
	}
	return nil
}
func (cb *Callbacks) OnFetchResponse(req *discoverygrpc.DiscoveryRequest, resp *discoverygrpc.DiscoveryResponse) {
	log.Infof("OnFetchResponse... Resquest[%v],  Response[%v]", req.TypeUrl, resp.TypeUrl)
}

func (cb *Callbacks) OnDeltaStreamClosed(id int64, node *corev3.Node) {
	log.Infof("OnDeltaStreamClosed... %v", id)
}

func (cb *Callbacks) OnDeltaStreamOpen(ctx context.Context, id int64, typ string) error {
	log.Infof("OnDeltaStreamOpen... %v  of type %s", id, typ)
	return nil
}

func (cb *Callbacks) OnStreamDeltaRequest(i int64, request *discoverygrpc.DeltaDiscoveryRequest) error {
	log.Infof("OnStreamDeltaRequest... %v  of type %s", i, request)
	return nil
}

func (cb *Callbacks) OnStreamDeltaResponse(i int64, request *discoverygrpc.DeltaDiscoveryRequest, response *discoverygrpc.DeltaDiscoveryResponse) {
	log.Infof("OnStreamDeltaResponse... %v  of type %s", i, request)
}

type Callbacks struct {
	Signal   chan struct{}
	Debug    bool
	Fetches  int
	Requests int
	mu       sync.Mutex
}
