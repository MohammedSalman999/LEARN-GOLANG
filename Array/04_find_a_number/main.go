package main

import "fmt"

func main() {
	li := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := 5
	fmt.Println(li)
	fmt.Println("number : ",n)
	fmt.Println("is exists : ",find(li,n))
}

func find(l []int, n int) bool {
	for _, val := range l {
		if n == val {
			return true
		}
	}
	return false
}