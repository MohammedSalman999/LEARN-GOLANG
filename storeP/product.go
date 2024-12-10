package main

import "fmt"

type prodcut struct {
	title    string
	price    money
	released timestamp
}

func (p prodcut) String() string {
	s := fmt.Sprintf("Title : %s , Price : %s , Relased : %s ", p.title, p.price, p.released)
	return s
}

func (p *prodcut) discount(ratio float64) {
	p.price *= money(1 - ratio)
}
