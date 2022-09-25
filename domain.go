package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

func query(host, domain string, task chan<- bool, result chan<- string) {
	topDomain := host + domain
	if !dns(topDomain) {
		result <- topDomain
	} else if !checkWhois(host, domain) {
		result <- topDomain
	}
	task <- true
}

func queryPrint(host, domain string, task chan<- bool) {
	topDomain := host + domain
	if !dns(topDomain) {
		fmt.Printf("Find Unregistry Domain: %s\n", topDomain)
	} else if !checkWhois(host, domain) {
		fmt.Printf("Find Unregistry Domain: %s\n", topDomain)
	}
	task <- true
}

func dns(domain string) bool {
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(1000), // 1s
			}
			return d.DialContext(ctx, "udp", "8.8.8.8:53")
		},
	}

	_, err := resolver.LookupHost(context.Background(), domain)
	if err != nil {
		fmt.Println(err)
		return false
		//} else {
		//	for _, address := range addresses {
		//		fmt.Printf("%s %s\n", domain, address)
		//	}
	}
	return true
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
