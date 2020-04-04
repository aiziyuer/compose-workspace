package test

import (
	"github.com/mwitkow/go-http-dialer"
	"golang.org/x/net/proxy"
	"gopkg.in/elazarl/goproxy.v1"
	"log"
	"net"
	"net/http"
	"net/url"
	"testing"
)

func TestHttpProxy2HttpUpstream(t *testing.T) {

	httpProxy := goproxy.NewProxyHttpServer()
	httpProxy.Verbose = true

	httpProxy.ConnectDial = func(network string, addr string) (net.Conn, error) {

		proxyUrl, _ := url.Parse("http://10.10.10.254:3128")
		d := http_dialer.New(proxyUrl)

		return d.Dial(network, addr)
	}

	log.Fatal(http.ListenAndServe(":8080", httpProxy))
}

func TestHttpProxy2SocketUpstream(t *testing.T) {

	httpProxy := goproxy.NewProxyHttpServer()
	httpProxy.Verbose = true

	httpProxy.ConnectDial = func(network string, addr string) (net.Conn, error) {

		d, _ := proxy.SOCKS5("tcp", "10.10.10.254:1080", nil, proxy.Direct)

		return d.Dial(network, addr)
	}

	log.Fatal(http.ListenAndServe(":8080", httpProxy))
}
