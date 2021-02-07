package config

import "fmt"

// Host keeps the details of a host service
type Host struct {
	// Address the address of a host
	Address string `yaml:"address"`
	// Port port on which the host listens
	Port int `yaml:"port"`
}

// FullAddress returns the full address of a host
func (h Host) FullAddress() string {
	return fmt.Sprintf("%s:%d", h.Address, h.Port)
}

// Service keeps the information of a service where the proxy forwards the requests
type Service struct {
	// Name the name of the service
	Name string `yaml:"name"`
	// Domain the full domain of the service
	Domain string `yaml:"domain"`
	// Hosts all the hosts to which the requests are forwards for this service
	Hosts []Host `yaml:"hosts"`
}

// Proxy keeps the entire proxy configuration
type Proxy struct {
	// Listen details related to the host on which the proxy listens for inbound requests
	Listen Host `yaml:"listen" json:"listen"`
	// Services all the services for which this proxy can forwards requests
	Services []Service `yaml:"services"`
}

// ProxyCfg container for proxy configuration
type ProxyCfg struct {
	Proxy Proxy `yaml:"proxy"`
}

// FindService looks up a service in the configuration by domain
func (p ProxyCfg) FindService(domain string) (Service, error) {
	for _, service := range p.Proxy.Services {
		if service.Domain == domain {
			return service, nil
		}
	}
	return Service{}, fmt.Errorf("no service found for domain %q", domain)
}

// ServiceHTTPScheme defines the HTTP scheme of a downstream service
const ServiceHTTPScheme = "http"
