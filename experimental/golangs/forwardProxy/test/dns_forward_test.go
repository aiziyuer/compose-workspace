package test

import (
	"github.com/gogf/gf/encoding/gparser"
	"github.com/miekg/dns"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"testing"
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
		r.SetRcode(req, dns.RcodeSuccess)

		for _, q := range req.Question {
			switch q.Qtype {
			default:
				r.SetRcode(req, dns.RcodeNameError)
			case dns.TypeA:
				a := &dns.A{
					Hdr: dns.RR_Header{Name: q.Name, Rrtype: q.Qtype, Class: dns.ClassINET, Ttl: 0},
					A:   net.ParseIP("0.0.0.0").To4(),
				}
				r.Answer = append(r.Answer, a)
			case dns.TypeAAAA:
				r.SetRcode(req, dns.RcodeNotImplemented)
			case dns.TypeDS:
				r.SetRcode(req, dns.RcodeNotImplemented)
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
