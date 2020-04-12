package test

import (
	"github.com/gogf/gf/encoding/gparser"
	"github.com/gogf/gf/util/gconv"
	"github.com/miekg/dns"
	"github.com/sensepost/godoh/dnsclient"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"testing"
	"time"
)

// 参考: https://ymmt2005.hatenablog.com/entry/2016/03/13/Transparent_SOCKS_proxy_in_Go_to_replace_NAT
func TestTCPDnsForward(t *testing.T) {

	h := dns.NewServeMux()
	h.HandleFunc(".", func(response dns.ResponseWriter, message *dns.Msg) {

		logrus.Info(gparser.MustToJsonString(message))

	})

	log.Fatal(dns.ListenAndServe("0.0.0.0:10053", "tcp", h))

}

func TestUDPDnsForward(t *testing.T) {

	h := dns.NewServeMux()

	// 测试方法:  dig @127.0.0.1 -p10053 www.google.com
	h.HandleFunc(".", func(rw dns.ResponseWriter, req *dns.Msg) {

		r := new(dns.Msg)
		r.SetReply(req)
		r.RecursionAvailable = req.RecursionDesired
		r.SetRcode(req, dns.RcodeSuccess)

		for _, q := range req.Question {
			switch q.Qtype {
			default:
				defaultResolver := &dns.Client{
					Net:          "tcp",
					ReadTimeout:  1000 * time.Second,
					WriteTimeout: 1000 * time.Second,
				}

				ret, _, err := defaultResolver.Exchange(
					new(dns.Msg).SetQuestion(q.Name, q.Qtype),
					"114.114.114.114:53",
				)
				// handle failed
				if err != nil {
					r.SetRcode(req, dns.RcodeServerFailure)
					logrus.Printf("Error: DNS:" + err.Error())
					continue
				}
				// domain not found
				if ret != nil && (ret.Rcode != dns.RcodeSuccess || len(ret.Answer) == 0) {
					r.SetRcode(req, dns.RcodeNameError)
					continue
				}
				r.Answer = append(r.Answer, ret.Answer[0])
			case dns.TypeA:
				// DoH
				client := &dnsclient.CloudflareDNS{BaseURL: "https://1.1.1.1/dns-query"}
				tmpResp := client.Lookup(q.Name, q.Qtype)
				logrus.Debugf("query from cloudflare: %s", gparser.MustToJsonString(tmpResp))

				a := &dns.A{
					Hdr: dns.RR_Header{Name: dns.Fqdn(q.Name), Rrtype: q.Qtype, Class: dns.ClassINET, Ttl: gconv.Uint32(tmpResp.TTL)},
					A:   net.ParseIP(tmpResp.Data).To4(),
				}
				r.Answer = append(r.Answer, a)
			case dns.TypeAAAA:
				// DoH
				client := &dnsclient.CloudflareDNS{BaseURL: "https://1.1.1.1/dns-query"}
				tmpResp := client.Lookup(q.Name, q.Qtype)
				logrus.Debugf("query from cloudflare: %s", gparser.MustToJsonString(tmpResp))

				a := &dns.AAAA{
					Hdr:  dns.RR_Header{Name: dns.Fqdn(q.Name), Rrtype: q.Qtype, Class: dns.ClassINET, Ttl: gconv.Uint32(tmpResp.TTL)},
					AAAA: net.ParseIP(tmpResp.Data),
				}
				r.Answer = append(r.Answer, a)
			}
		}

		err := rw.WriteMsg(r)
		if err != nil {
			logrus.Warnf("Error: Writing Response:%v\n", err)
		}
		_ = rw.Close()

	})

	log.Fatal(dns.ListenAndServe("0.0.0.0:10053", "udp", h))

}
