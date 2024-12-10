package main

import "fmt"

func main() {
	store := list{
		{title: "harry potter", price: 250, released: toTimeStamp("jan-2006")},
	}
	store.discount(10)
	fmt.Println(store)
}
