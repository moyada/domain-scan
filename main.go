package main

import (
	"flag"
	"fmt"
	//"sort"
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

	//size := len(enWords) * len(enWords)
	task := make(chan bool, 512)
	size := 0
	//result := make(chan string, size)

	g := New(*s)
	for {
		host := g.Next()
		if host == "" {
			break
		}
		size++
		go queryPrint(prefix+host, domain, task)
	}

	for i := 0; i < size; i++ {
		<-task
	}

	//l := len(result)
	//r := make([]string, 0, l)
	//for i := 0; i < l; i++ {
	//	r = append(r, <-result)
	//}
	//
	//sort.Strings(r)
	//for _, s := range r {
	//	fmt.Println(s)
	//}

	fmt.Println("domain scan completed")
}

func tes() {
	topDomain := "angs.com"
	if !dns(topDomain) {
		fmt.Println("dns:", topDomain)
	} else if !checkWhois("angs", ".com") {
		fmt.Println("dns:", topDomain)
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
