package main

import "fmt"

var (
	enWords = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
)

type Generator struct {
	prefix []byte
	index  int
}

func New(s int) *Generator {
	bs := make([]byte, s-1)
	for i := 0; i < s-1; i++ {
		bs[i] = enWords[0]
	}
	return &Generator{prefix: bs, index: -1}
}

func (g *Generator) Next() string {
	n := g.index + 1
	if n == len(enWords) {
		if !g.growth(len(g.prefix) - 1) {
			return ""
		}
		g.index = -1
		return g.Next()
	}
	e := enWords[n]
	g.index = n
	return fmt.Sprintf("%s%c", g.prefix, e)
}

func (g *Generator) growth(level int) bool {
	if level < 0 {
		return false
	}
	b := g.prefix[level]

	nd := -1
	for i, word := range enWords {
		if word == b {
			nd = i + 1
			break
		}
	}

	if nd == len(enWords) {
		g.prefix[level] = enWords[0]
		return g.growth(level - 1)
	}

	g.prefix[level] = enWords[nd]
	return true
}
