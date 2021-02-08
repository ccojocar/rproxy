package cmd

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ccojocar/rproxy/pkg/config"
	"github.com/ccojocar/rproxy/pkg/lb"
	"github.com/ccojocar/rproxy/pkg/proxy"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Starts the reverse proxy",
	RunE: func(cmd *cobra.Command, args []string) error {
		http.HandleFunc("/", handleProxy)
		http.HandleFunc("/health", handleHealth)
		if err := http.ListenAndServe(listenAddress(), nil); err != nil {
			return fmt.Errorf("listening on %s: %w", listenAddress(), err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

// listenAddress returns the proxy listen address as provided in the configuration
func listenAddress() string {
	address := proxyConfig.Proxy.Listen.Address
	port := proxyConfig.Proxy.Listen.Port
	return fmt.Sprintf("%s:%d", address, port)
}

// handleProxy main entry point for handling all proxy requests
func handleProxy(res http.ResponseWriter, req *http.Request) {
	domain := req.Host
	if domain == "" {
		log.Warnf("Not service found in the %q HTTP header", "Host")
		http.Error(res, "Not service provided", http.StatusBadRequest)
		return
	}
	svc, err := proxyConfig.FindService(domain)
	if err != nil {
		log.Warnf("Service not found for domain %q", domain)
		http.Error(res, "Service not found", http.StatusBadRequest)
		return
	}
	rlb := &lb.RandomLoadBalancer{}
	host, err := rlb.GetHost(svc)
	if err != nil {
		log.Warnf("Failed to load balance the %q service: %s", svc.Name, err)
		http.Error(res, "Service not found", http.StatusBadRequest)
		return
	}
	target := &url.URL{
		Scheme: config.ServiceHTTPScheme,
		Host:   host.FullAddress(),
	}
	proxy.ServeHTTP(target, res, req)
}

// handleHealth health probe
func handleHealth(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
}
