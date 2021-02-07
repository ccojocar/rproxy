package lb

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/ccojocar/rproxy/pkg/config"
)

// RandomLoadBalancer implements a random load balancing strategy
type RandomLoadBalancer struct {
}

// GetHost returns a random host from the given service configuration
func (r *RandomLoadBalancer) GetHost(svc config.Service) (config.Host, error) {
	max := len(svc.Hosts)
	if max == 0 {
		return config.Host{}, fmt.Errorf("no hosts found in the %q service configuration", svc.Name)
	}
	index, err := getHostIndex(int64(max))
	if err != nil {
		return config.Host{}, err
	}
	return svc.Hosts[index], nil
}

// getHostIndex generate a random service index within [0, max)
func getHostIndex(max int64) (int64, error) {
	i, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return i.Int64(), fmt.Errorf("generate random proxy service index: %w", err)
	}
	return i.Int64(), nil
}
