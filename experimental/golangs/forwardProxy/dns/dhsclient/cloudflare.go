package dhsclient

import (
	"encoding/json"
	"github.com/miekg/dns"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type CloudflareDNS struct {
	BaseURL string
}

func (c *CloudflareDNS) Lookup(name string, rType uint16) RequestResponse {

	client := http.Client{
		Timeout: time.Second * 20,
	}

	req, err := http.NewRequest("GET", c.BaseURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("accept", "application/dns-json")

	q := req.URL.Query()
	q.Add("name", name)
	q.Add("type", dns.TypeToString[rType])
	q.Add("cd", "false") // ignore DNSSEC
	q.Add("do", "false") // ignore DNSSEC
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	resp := RequestResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Fatal(err)
	}

	return resp
}
