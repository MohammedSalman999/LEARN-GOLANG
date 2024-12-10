package main

import "fmt"

type product struct {
	title    string
	price    money
	released float64
}

func (p product) String() string {
	return fmt.Sprintf("name %s , price : %s released : %f", p.title, p.price, p.released)
}

func (p product) discount(ratio float64) {
	p.price *= money(1 - ratio)
}
