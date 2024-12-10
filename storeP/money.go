package main

import "fmt"

type money float64

func (m money) String() string {
	s := fmt.Sprintf("$%.2f", m)
	return s
}
