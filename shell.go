package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

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
