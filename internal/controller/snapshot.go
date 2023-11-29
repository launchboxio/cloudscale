package controller

import (
	clusterv3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpointv3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	listenerv3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	routev3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	routerv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"
	hcmv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	cachev3 "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	resourcev3 "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/launchboxio/cloudscale/internal/api"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"time"
)

type SnapshotInfo struct {
	Version                string
	Listeners              []*api.Listener
	TargetGroups           []*api.TargetGroup
	TargetGroupAttachments []*api.TargetGroupAttachment
	Certificates           []*api.Certificate
}

func toClusters(targetGroups []*api.TargetGroup, attachments []*api.TargetGroupAttachment) []types.Resource {
	var result []types.Resource
	for _, tg := range targetGroups {
		targetGroupAttachments := getTargetGroupAttachments(tg.Id, attachments)
		result = append(result, toCluster(tg, targetGroupAttachments))
	}
	return result
}

// TODO: Fairly ineffecient. We should maybe prebuild the attachment mapping
func getTargetGroupAttachments(targetGroupId string, attachments []*api.TargetGroupAttachment) []*api.TargetGroupAttachment {
	var result []*api.TargetGroupAttachment
	for _, attachment := range attachments {
		if attachment.TargetGroupId != targetGroupId {
			continue
		}
		result = append(result, attachment)
	}
	return result
}

func toCluster(targetGroup *api.TargetGroup, attachments []*api.TargetGroupAttachment) *clusterv3.Cluster {
	return &clusterv3.Cluster{
		Name:                 targetGroup.Name,
		ConnectTimeout:       durationpb.New(5 * time.Second),
		ClusterDiscoveryType: &clusterv3.Cluster_Type{Type: clusterv3.Cluster_LOGICAL_DNS},
		LbPolicy:             clusterv3.Cluster_ROUND_ROBIN,
		LoadAssignment:       toEndpoints(targetGroup, attachments),
		DnsLookupFamily:      clusterv3.Cluster_V4_ONLY,
	}
}

func toEndpoints(targetGroup *api.TargetGroup, attachments []*api.TargetGroupAttachment) *endpointv3.ClusterLoadAssignment {
	var lbEndpoints []*endpointv3.LbEndpoint
	for _, attachment := range attachments {
		lbEndpoint := &endpointv3.LbEndpoint{
			HostIdentifier: &endpointv3.LbEndpoint_Endpoint{
				Endpoint: &endpointv3.Endpoint{
					Address: &corev3.Address{
						Address: &corev3.Address_SocketAddress{
							SocketAddress: &corev3.SocketAddress{
								Protocol: corev3.SocketAddress_TCP,
								Address:  attachment.IpAddress.String(),
								PortSpecifier: &corev3.SocketAddress_PortValue{
									PortValue: uint32(attachment.Port),
								},
							},
						},
					},
				},
			},
		}
		lbEndpoints = append(lbEndpoints, lbEndpoint)
	}
	return &endpointv3.ClusterLoadAssignment{
		ClusterName: targetGroup.Name,
		Endpoints: []*endpointv3.LocalityLbEndpoints{{
			LbEndpoints: lbEndpoints,
		}},
	}
}

func toRoutes() []types.Resource {
	return []types.Resource{toRoute()}
}

func toRoute() *routev3.RouteConfiguration {
	return &routev3.RouteConfiguration{
		Name: "local_route",
		VirtualHosts: []*routev3.VirtualHost{{
			Name:    "local_service",
			Domains: []string{"*"},
			Routes: []*routev3.Route{{
				Match: &routev3.RouteMatch{
					PathSpecifier: &routev3.RouteMatch_Prefix{
						Prefix: "/",
					},
				},
				Action: &routev3.Route_Route{
					Route: &routev3.RouteAction{
						ClusterSpecifier: &routev3.RouteAction_Cluster{
							Cluster: "cluster-0",
						},
						HostRewriteSpecifier: &routev3.RouteAction_HostRewriteLiteral{
							HostRewriteLiteral: "127.0.0.1",
						},
					},
				},
			}},
		}},
	}
}

func toListeners(listeners []*api.Listener) []types.Resource {
	var result []types.Resource
	for _, listener := range listeners {
		result = append(result, toHttpListener(listener))
	}
	return result
}

func toHttpListener(listener *api.Listener) *listenerv3.Listener {
	routerConfig, _ := anypb.New(&routerv3.Router{})
	// HTTP filter configuration
	manager := &hcmv3.HttpConnectionManager{
		CodecType:  hcmv3.HttpConnectionManager_AUTO,
		StatPrefix: "http",
		RouteSpecifier: &hcmv3.HttpConnectionManager_Rds{
			Rds: &hcmv3.Rds{
				ConfigSource: &corev3.ConfigSource{
					ResourceApiVersion: resourcev3.DefaultAPIVersion,
					ConfigSourceSpecifier: &corev3.ConfigSource_ApiConfigSource{
						ApiConfigSource: &corev3.ApiConfigSource{
							TransportApiVersion:       resourcev3.DefaultAPIVersion,
							ApiType:                   corev3.ApiConfigSource_GRPC,
							SetNodeOnFirstMessageOnly: true,
							GrpcServices: []*corev3.GrpcService{{
								TargetSpecifier: &corev3.GrpcService_EnvoyGrpc_{
									EnvoyGrpc: &corev3.GrpcService_EnvoyGrpc{ClusterName: "xds_cluster"},
								},
							}},
						},
					},
				},

				RouteConfigName: "local_route",
			},
		},
		HttpFilters: []*hcmv3.HttpFilter{{
			Name:       "http-router",
			ConfigType: &hcmv3.HttpFilter_TypedConfig{TypedConfig: routerConfig},
		}},
	}
	pbst, err := anypb.New(manager)
	if err != nil {
		panic(err)
	}

	return &listenerv3.Listener{
		Name: listener.Name,
		Address: &corev3.Address{
			Address: &corev3.Address_SocketAddress{
				SocketAddress: &corev3.SocketAddress{
					Protocol: corev3.SocketAddress_TCP,
					Address:  listener.IpAddress.String(),
					PortSpecifier: &corev3.SocketAddress_PortValue{
						PortValue: uint32(listener.Port),
					},
				},
			},
		},
		FilterChains: []*listenerv3.FilterChain{{
			Filters: []*listenerv3.Filter{{
				Name: "http-connection-manager",
				ConfigType: &listenerv3.Filter_TypedConfig{
					TypedConfig: pbst,
				},
			}},
		}},
	}
}

func generateSnapshot(info *SnapshotInfo) (*cachev3.Snapshot, error) {
	return cachev3.NewSnapshot(info.Version,
		map[resourcev3.Type][]types.Resource{
			resourcev3.ClusterType:  toClusters(info.TargetGroups, info.TargetGroupAttachments),
			resourcev3.RouteType:    toRoutes(),
			resourcev3.ListenerType: toListeners(info.Listeners),
		},
	)
}

func buildSnapshot(svc *api.Service) (*cachev3.Snapshot, error) {

	certificates, err := svc.ListCertificates()
	if err != nil {
		return nil, err
	}

	listeners, err := svc.ListListeners()
	if err != nil {
		return nil, err
	}

	targetGroups, err := svc.ListTargetGroups()
	if err != nil {
		return nil, err
	}

	info := &SnapshotInfo{
		Version:      "1",
		TargetGroups: targetGroups,
		Certificates: certificates,
		Listeners:    listeners,
	}

	return generateSnapshot(info)
}
