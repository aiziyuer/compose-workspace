package dhsclient

import "github.com/miekg/dns"

// Client is an interface all clients should conform to.
type Client interface {
	Lookup(name string, rType uint16) *dns.Msg
	LookupAppend(r *dns.Msg, name string, rType uint16)
}

func NewCloudFlareDNS(modOptions ...ModDoHOption) *CloudflareDNS {

	option := DoHOption{BaseURL: "https://cloudflare-dns.com/dns-query"}
	for _, fn := range modOptions {
		fn(&option)
	}

	return &CloudflareDNS{BaseURL: option.BaseURL}
}
