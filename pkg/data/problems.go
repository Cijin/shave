package data

import (
	"fmt"
	"strings"
)

// a single field can have multiple things wrong with it.
// It's a better user experience to see them at once
// rather than re-submitting the form
type Problems map[string][]string

func NewProblems() Problems {
	return make(map[string][]string)
}

func (p Problems) Add(field string, msg string) {
	if _, ok := p[field]; !ok {
		p[field] = []string{}
	}
	p[field] = append(p[field], msg)
}

func (p Problems) AddAll(field string, msgs []string) {
	if _, ok := p[field]; !ok {
		p[field] = []string{}
	}
	p[field] = msgs
}

func (p Problems) Any() bool {
	return len(p) > 0
}

func (p Problems) Get(field string) []string {
	return p[field]
}

func (p Problems) Has(field string) bool {
	return len(p[field]) > 0
}

func (p Problems) String() string {
	var b strings.Builder

	for k, v := range p {
		fmt.Fprintf(&b, "%s:%s\n", k, strings.Join(v, "."))
	}

	return b.String()
}
