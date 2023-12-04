package controller

import (
	"context"
	"fmt"
	serverv3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"github.com/launchboxio/cloudscale/internal/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net"

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
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	svc := &api.Service{Db: db}
	if err = svc.Init(); err != nil {
		return err
	}
	ctx := context.Background()
	signal := make(chan struct{})
	event := make(chan struct{})
	cb := &Callbacks{
		Signal:   signal,
		Fetches:  0,
		Requests: 0,
	}

	cache = cachev3.NewSnapshotCache(true, cachev3.IDHash{}, nil)
	srv := serverv3.NewServer(ctx, cache, cb)

	go func() {
		if err := c.runApiServer(svc, event, cache); err != nil {
			log.WithError(err).Error("Failed starting API server")
		}
	}()

	go c.runSnapshotHandler(ctx, svc, event)

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
		reflection.Register(grpcServer)
		if err = grpcServer.Serve(lis); err != nil {
			log.WithError(err).Error("Failed running GRPC server")
		}
	}()
	<-ctx.Done()

	grpcServer.GracefulStop()
}

func (c *Controller) runApiServer(svc *api.Service, channel chan struct{}, snapshotCache cachev3.SnapshotCache) error {
	srv := api.New(svc, channel, snapshotCache)
	return srv.Run(c.options.HttpAddress)
}

func (c *Controller) runSnapshotHandler(ctx context.Context, svc *api.Service, channel chan struct{}) {
	log.Info("Generating initial snapshot")
	if err := handleSnapshotEvent(svc); err != nil {
		log.WithError(err).Error("Failed generating snapshot")
	}
	for {
		// check on channel
		select {
		case <-channel:
			if err := handleSnapshotEvent(svc); err != nil {
				log.WithError(err).Error("Failed generating snapshot")
				continue
			}
			log.Info("New Snapshot saved successfully")
		case <-ctx.Done():
			return
		default:
		}
		// continue working
	}
}

func handleSnapshotEvent(svc *api.Service) error {
	snap, err := buildSnapshot(svc)
	if err != nil {
		fmt.Println("Failed building snapshot")
		return err
	}
	if err := snap.Consistent(); err != nil {
		fmt.Println("Snapshot not consistent")
		return err
	}
	return cache.SetSnapshot(context.Background(), "test-id", snap)
}
