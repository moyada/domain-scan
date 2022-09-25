package main

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"os/exec"
	"strings"
	"time"
)

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
		//fmt.Println(err)
		return -1
	}

	//for _, address := range addresses {
	//	fmt.Printf("%s %s\n", domain, address)
	//}
	return len(addresses)
}

func whois(domain string) bool {
	cmd := exec.Command("whois", domain)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("whois error:", domain)
		return false
	}
	out := stdout.String()
	return !strings.Contains(out, "No match for domain")
}
