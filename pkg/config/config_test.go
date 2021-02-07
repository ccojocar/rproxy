package config_test

import (
	"testing"

	"github.com/ccojocar/rproxy/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestFindService(t *testing.T) {
	cfg := config.ProxyCfg{
		Proxy: config.Proxy{
			Services: []config.Service{
				{
					Name:   "test1",
					Domain: "test1.my-company.org",
					Hosts: []config.Host{
						{
							Address: "test1",
							Port:    1234,
						},
					},
				},
				{
					Name:   "test2",
					Domain: "test2.my-company.org",
					Hosts: []config.Host{
						{
							Address: "test2",
							Port:    1234,
						},
					},
				},
			},
		},
	}

	svc, err := cfg.FindService("test1.my-company.org")
	assert.NoError(t, err)
	assert.Equal(t, "test1.my-company.org", svc.Domain)

	_, err = cfg.FindService("test3.company.org")
	assert.Error(t, err)
}
