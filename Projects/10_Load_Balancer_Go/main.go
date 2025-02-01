package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	utils "github.com/devilzs1/load-balancer-go/utils"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(w http.ResponseWriter, r *http.Request)
	HealthCheck() bool
}

type myServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func newServer(addr string) *myServer {
	serverUrl, err := url.Parse(addr)

	utils.HandleError(err)

	return &myServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

type LoadBalancer struct {
	port            string
	RoundRobinCount int
	servers         []Server
}

func NewLoadBalancer(port string, server []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		RoundRobinCount: 0,
		servers:         server,
	}
}

func (lb *LoadBalancer) getNextAvailableServer() Server {
	for {
		server := lb.servers[lb.RoundRobinCount]
		if server.IsAlive() {
			lb.RoundRobinCount = (lb.RoundRobinCount + 1) % len(lb.servers)
			return server
		}
		lb.RoundRobinCount = (lb.RoundRobinCount + 1) % len(lb.servers)
	}
}

func (lb *LoadBalancer) serveProxyServer(w http.ResponseWriter, r *http.Request) {
	server := lb.getNextAvailableServer()
	fmt.Printf("Forwarding request to server : %s\n", server.Address())
	server.Serve(w, r)
}

func (s *myServer) Address() string {
	return s.addr
}

func (s *myServer) IsAlive() bool {
	return true
}

func (s *myServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

func main() {
	fmt.Println("Learning Load Balancer Implementation using GO Programming")

	servers := []Server{
		myServer("https://google.com"),
		myServer("https://bing.com"),
		myServer("https://facebook.com"),
	}

	lb := NewLoadBalancer("4000", servers)
	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		lb.serveProxyServer(w, r)
	}

	http.HandleFunc("/", handleRedirect)

	http.ListenAndServe(":"+lb.port, nil)
	fmt.Printf("Server is listening at http://localhost:%s\n", lb.port)
}
