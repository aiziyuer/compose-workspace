package test

import (
	"github.com/armon/go-socks5"
	"github.com/mwitkow/go-http-dialer"
	"golang.org/x/net/context"
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

	// curl -vI -x 127.0.0.1:8080 https://www.google.com
	log.Fatal(http.ListenAndServe(":8080", httpProxy))
}

func TestHttpProxy2SocketUpstream(t *testing.T) {

	httpProxy := goproxy.NewProxyHttpServer()
	httpProxy.Verbose = true

	httpProxy.ConnectDial = func(network string, addr string) (net.Conn, error) {

		d, _ := proxy.SOCKS5("tcp", "10.10.10.254:1080", nil, proxy.Direct)

		return d.Dial(network, addr)
	}

	// curl -vI -x 127.0.0.1:8080 https://www.google.com
	log.Fatal(http.ListenAndServe(":8080", httpProxy))
}

func TestSocketProxy2HttpUpstream(t *testing.T) {

	conf := &socks5.Config{
		Dial: func(ctx context.Context, network, addr string) (conn net.Conn, err error) {

			proxyUrl, _ := url.Parse("http://10.10.10.254:3128")
			d := http_dialer.New(proxyUrl)

			return d.Dial(network, addr)
		},
	}
	server, _ := socks5.New(conf)

	// curl -vI -x socks5h://127.0.0.1:8080 https://www.google.com
	log.Fatal(server.ListenAndServe("tcp", ":8080"))
}

func TestSocketProxy2SocketUpstream(t *testing.T) {

	conf := &socks5.Config{
		Dial: func(ctx context.Context, network, addr string) (conn net.Conn, err error) {

			d, _ := proxy.SOCKS5("tcp", "10.10.10.254:1080", nil, proxy.Direct)

			return d.Dial(network, addr)
		},
	}
	server, _ := socks5.New(conf)

	// curl -vI -x socks5h://127.0.0.1:8080 https://www.google.com
	log.Fatal(server.ListenAndServe("tcp", ":8080"))
}
