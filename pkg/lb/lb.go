package lb

import "github.com/ccojocar/rproxy/pkg/config"

// LoadBalancer defines an interface which should be implemented by the load balancer strategy
type LoadBalancer interface {
	// GetHost returns a free host based on the load balancing strategy
	GetHost(svc config.Service) (config.Host, error)
}
