package controller

import (
	"context"
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	serverv3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"github.com/launchboxio/cloudscale/internal/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"sync"

	bolt "go.etcd.io/bbolt"

	discoverygrpc "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	cachev3 "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
)

const (
	XdsCluster = "xds_cluster"
	Ads        = "ads"
	Xds        = "xds"
	Rest       = "rest"
)

var (
	cache cachev3.SnapshotCache
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

type Options struct {
	GrpcAddress string
	HttpAddress string
}

func New(opts *Options) *Controller {
	return &Controller{
		options: opts,
	}
}

type Controller struct {
	options *Options
}

// Run will start the controller service, which exposes
// an xDS endpoint for Envoy configuration. It will also
// start the API server for handling external configurations
func (c *Controller) Run() error {

	// Open our database connection
	db, err := bolt.Open("./bolt", 0600, nil)
	if err != nil {
		return err
	}

	// Ensure our buckets exist
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, b := range []string{
		api.ListenersBucket,
		api.CertificateBucket,
	} {
		log.Infof("Ensuring bucket %s exists", b)
		if _, err := tx.CreateBucketIfNotExists([]byte(b)); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	defer db.Close()

	ctx := context.Background()
	signal := make(chan struct{})
	cb := &Callbacks{
		Signal:   signal,
		Fetches:  0,
		Requests: 0,
	}

	cache = cachev3.NewSnapshotCache(true, cachev3.IDHash{}, nil)
	srv := serverv3.NewServer(ctx, cache, cb)

	go func() {
		if err := c.runApiServer(db); err != nil {
			log.WithError(err).Error("Failed starting API server")
		}
	}()
	c.runXdsServer(ctx, srv)

	return nil
}

func (c *Controller) runXdsServer(ctx context.Context, srv serverv3.Server) {
	var grpcOptions []grpc.ServerOption
	grpcServer := grpc.NewServer(grpcOptions...)

	lis, err := net.Listen("tcp", c.options.GrpcAddress)
	if err != nil {
		log.WithError(err).Fatal("Failed listening to port")
		return
	}

	discoverygrpc.RegisterAggregatedDiscoveryServiceServer(grpcServer, srv)

	log.Info("Management server listening")

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.WithError(err).Error("Failed running GRPC server")
		}
	}()
	<-ctx.Done()

	grpcServer.GracefulStop()
}

func (c *Controller) runApiServer(db *bolt.DB) error {
	srv := api.New(db)
	return srv.Run(c.options.HttpAddress)
}
