package main

import "strings"

type list []*prodcut

func (li list) String() string {
	if len(li) == 0 {
		return "waiting for the list"
	}

	var str strings.Builder
	for _, p := range li {
		str.WriteString("* ")
		str.WriteString(p.String())
		str.WriteRune('\n')
	}
	return str.String()
}

func (li list) discount(ratio float64) {
	for _, p := range li {
		p.discount(ratio)
	}
}
