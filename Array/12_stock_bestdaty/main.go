// find the best day to buy the stock and sell the stock
package main

import (
	"fmt"
)

type money int

func (m money) String() string {
	return fmt.Sprintf("$%d", m)
}

func finDay(slice []int) (money, money, money, money, money) {
	minp := slice[0] // minimum price
	maxp := 0        // maximum price
	buyd := 0
	selld := 0
	profit := 0

	for day, price := range slice {
		if price < minp {
			minp = price
			buyd = day
		} else if profit = price - minp; profit > maxp {
			maxp = profit
			selld = day
		}
	}

	return money(buyd), money(maxp), money(selld), money(slice[selld]), money(profit)
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	buyDay, buyprice, sellday, sellprice, profit := finDay(slice)
	fmt.Println("given stock list", slice)
	fmt.Printf("final decision : \n\t\t buy day : \t%d\t buyprice : \t%s \n\t\t sell day: \t%d\t sellprice : \t%s\t\n\t\t profit : \t%d", buyDay+1, buyprice, sellday+1, sellprice, profit)
}
