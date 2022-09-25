package main

import "fmt"

type Callback interface {
	Accept(result string)
}

type Console struct {
}

func (c *Console) Accept(result string) {
	fmt.Printf("Find Unregistry Domain: %s\n", result)
}
