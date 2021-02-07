package lb_test

import (
	"testing"

	"github.com/ccojocar/rproxy/pkg/config"
	"github.com/ccojocar/rproxy/pkg/lb"
	"github.com/stretchr/testify/assert"
)

func TestRandomLoadBalancer(t *testing.T) {
	rlb := &lb.RandomLoadBalancer{}

	// Test for configuration with some hosts
	svc := config.Service{
		Name:   "test",
		Domain: "test.my-company.org",
		Hosts: []config.Host{
			{
				Address: "test1",
				Port:    1,
			},
			{
				Address: "test2",
				Port:    2,
			},
			{
				Address: "test3",
				Port:    3,
			},
		},
	}

	host, err := rlb.GetHost(svc)
	assert.NoError(t, err)
	assert.NotEmpty(t, host)

	// Test the empty configuration
	svc = config.Service{}
	host, err = rlb.GetHost(svc)
	assert.Error(t, err)
	assert.Empty(t, host)
}
