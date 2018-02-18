package ovc

import (
	"io"

	"github.com/golang/glog"
	"k8s.io/kubernetes/pkg/cloudprovider"
	"k8s.io/kubernetes/pkg/controller"
)

const (
	providerName = "ovc"
)

func init() {
	cloudprovider.RegisterCloudProvider(providerName, func(config io.Reader) (cloudprovider.Interface, error) {
		return newOVCProvider(config)
	})
}

func newOVCProvider(config io.Reader) (cloudprovider.Interface, error) {
	glog.Info("creating ovc provider")
	return &OVC{}, nil
}

//OVC Cloud interface
type OVC struct {
}

//Initialize cloud
func (o *OVC) Initialize(clientBuilder controller.ControllerClientBuilder) {
	glog.Info("initializing ovc provider")
}

//LoadBalancer gets a load balancer
func (o *OVC) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	return nil, false
}

//Instances gets instances
func (o *OVC) Instances() (cloudprovider.Instances, bool) {
	return nil, false
}

//Zones gets zones
func (o *OVC) Zones() (cloudprovider.Zones, bool) {
	return nil, false
}

//Clusters gets clusters
func (o *OVC) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false
}

//Routes gets routes
func (o *OVC) Routes() (cloudprovider.Routes, bool) {
	return nil, false
}

//ProviderName gets provider name
func (o *OVC) ProviderName() string {
	return providerName
}

//HasClusterID gets if has a cluster id
func (o *OVC) HasClusterID() bool {
	return true
}

//deprecated
func (o *OVC) ScrubDNS(nameservers, searches []string) (nsOut, srchOut []string) {
	return nil, nil
}
