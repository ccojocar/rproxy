package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// ServeHTTP reverse proxy a request to the target URL
// Note: this is a non-blocking call which uses a goroutine to handle
// the proxy connection with the client
func ServeHTTP(target *url.URL, res http.ResponseWriter, req *http.Request) {
	req.URL.Host = target.Host
	req.URL.Scheme = target.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = target.Host

	p := httputil.NewSingleHostReverseProxy(target)
	p.ServeHTTP(res, req)
}
