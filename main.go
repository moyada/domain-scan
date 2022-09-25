package main

import (
	"flag"
	"fmt"
	//"sort"
)

const (
	concurrency = 32
)

func main() {
	p := flag.String("p", "", "")
	d := flag.String("d", ".com", "")
	s := flag.Int("s", 0, "")
	flag.Parse()

	if *s == 0 {
		return
	}

	prefix := *p
	domain := *d

	task := make(chan bool, concurrency)
	for i := 0; i < concurrency; i++ {
		task <- true
	}

	g := New(*s)
	c := Console{}

	for {
		host := g.Next()
		if host == "" {
			break
		}
		<-task
		go queryPrint(prefix+host, domain, &c, task)
	}

	for i := 0; i < concurrency; i++ {
		<-task
	}

	fmt.Println("Domain Scan Completed.")
}

func main1() {
	topDomain := "angs.com"
	count := dns(topDomain)
	fmt.Println(count)
	switch {
	case count > 0:
		fmt.Println(topDomain)
	case count == -1:
		if !checkWhois("angs", ".com") {
			fmt.Println(topDomain)
		}
	}
}

func test() {
	g := New(4)
	for {
		s := g.Next()
		if s == "" {
			break
		}
		fmt.Println(s)
	}
}
