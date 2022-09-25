package main

import (
	"github.com/openrdap/rdap/bootstrap"
	"net/http"
	"time"
)

var (
	comApi = rdapApi("google.com")
	netApi = rdapApi("create.net")
)

func rdapApi(domain string) string {
	question := &bootstrap.Question{
		RegistryType: bootstrap.DNS,
		Query:        domain,
	}

	b := &bootstrap.Client{}

	var answer *bootstrap.Answer
	answer, err := b.Lookup(question)

	if err == nil {
		for _, url := range answer.URLs {
			return url.String() + "/domain/"
		}
	}
	return ""
}

func rdapQuery(host, domain string) bool {
	client := http.Client{Timeout: 10 * time.Second}

	var api string
	switch domain {
	case ".com":
		api = comApi
	case ".net":
		api = netApi
	}

	resp, err := client.Get(api + host + domain)
	if err != nil {
		return whois(host + domain)
	}
	return resp.StatusCode == http.StatusOK
}
