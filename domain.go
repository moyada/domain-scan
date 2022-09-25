package main

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
