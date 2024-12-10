package main

import "fmt"

func min(list []int) int {
	min := list[0]
	for _, val := range list {
		if val < min {
			min = val
		}
	}

	return min
}

func main() {
	list := []int{1, 2, 3, 4, 5, -1, 6}
	fmt.Println("list : \t ", list)
	fmt.Println("min value : \t ", min(list))
}
