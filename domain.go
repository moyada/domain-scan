package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

const (
	dnsServer = "8.8.8.8:53"
)

//func query(host, domain string, task chan<- bool, result chan<- string) {
//	topDomain := host + domain
//	if !dns(topDomain) {
//		result <- topDomain
//	} else if !checkWhois(host, domain) {
//		result <- topDomain
//	}
//	task <- true
//}

func queryPrint(host, domain string, cb Callback, task chan<- bool) {
	topDomain := host + domain

	count := dns(topDomain)
	switch {
	case count == 0:
		cb.Accept(topDomain)
	case count == -1:
		if !checkWhois(host, domain) {
			cb.Accept(topDomain)
		}
	}

	task <- true
}

func dns(domain string) int {
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(1000), // 1s
			}
			return d.DialContext(ctx, "udp", dnsServer)
		},
	}

	addresses, err := resolver.LookupHost(context.Background(), domain)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	//for _, address := range addresses {
	//	fmt.Printf("%s %s\n", domain, address)
	//}
	return len(addresses)
}

func checkWhois(host, domain string) bool {
	switch domain {
	case ".com":
		return rdapQuery(host, domain)
	case ".net":
		return rdapQuery(host, domain)
	default:
		return whois(host + domain)
	}
}
