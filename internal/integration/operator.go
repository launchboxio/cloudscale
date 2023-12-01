package integration

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"time"
)

const (
	// The annotation base
	CloudScaleBaseAnnotation string = "service.cloudscale.io"
	// SourceRangeAnnotation lists the IPs this service accepts traffic from
	SourceRangeAnnotation string = "load-balancer-source-ranges"
	// SslCertificateId specifies the SSL certificate to use for termination
	SslCertificateId string = "ssl-certificate"
	// ListenerIpAddress specifies the IP address of the remote load balancer
	ListenerIpAddress string = "listener-ip-address"
)

type Operator struct {
	clientset      *kubernetes.Clientset
	cloudscaleAddr string
}

func New(clientset *kubernetes.Clientset, cloudscaleAddr string) *Operator {
	return &Operator{clientset, cloudscaleAddr}
}

func (o *Operator) Create(obj interface{}) {
	fmt.Printf("service added: %s \n", obj)
}

func (o *Operator) Update(oldObj, newObj interface{}) {
	fmt.Printf("service changed \n")
}

func (o *Operator) Delete(obj interface{}) {
	fmt.Printf("service deleted: %s \n", obj)
}

// updateService will propagate the information returned from cloudscale
func (o *Operator) updateService(ctx context.Context, ipAddr string, port int32, svc *v1.Service) error {
	if len(svc.Status.LoadBalancer.Ingress) != 1 ||
		svc.Status.LoadBalancer.Ingress[0].IP != ipAddr ||
		len(svc.Status.LoadBalancer.Ingress[0].Ports) != 1 ||
		svc.Status.LoadBalancer.Ingress[0].Ports[0].Port != port {
		//svcOld := svc.DeepCopy()
		svc.Status.LoadBalancer.Ingress = []v1.LoadBalancerIngress{
			{
				IP: ipAddr,
				Ports: []v1.PortStatus{
					{Port: port},
				},
			},
		}

		_, err := o.clientset.CoreV1().Services(svc.Namespace).UpdateStatus(ctx, svc, metav1.UpdateOptions{})
		return err
	}
	return nil
}

func (o *Operator) ResourceEventHandlerFuncs() *cache.ResourceEventHandlerFuncs {
	return &cache.ResourceEventHandlerFuncs{
		AddFunc:    o.Create,
		DeleteFunc: o.Delete,
		UpdateFunc: o.Update,
	}
}

func (o *Operator) Run(ctx context.Context) error {
	watchlist := cache.NewListWatchFromClient(
		o.clientset.CoreV1().RESTClient(),
		string(v1.ResourceServices),
		v1.NamespaceAll,
		fields.Everything(),
	)
	_, controller := cache.NewInformer(
		watchlist,
		&v1.Service{},
		0,
		o.ResourceEventHandlerFuncs(),
	)

	stop := make(chan struct{})
	defer close(stop)
	go controller.Run(stop)

	for {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			return nil
		}
	}
}
