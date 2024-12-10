package main

import "fmt"

func min(list []int) int {
	max := list[0]
	smax := 0
	for _, val := range list {
		if val > max {
			smax = max
			max = val
		} else if val > smax && val != max {
			smax = val
		}
	}

	return smax
}

func main() {
	list := []int{1, 2, 3, 4, 5, -1, 6}
	fmt.Println("list : \t ", list)
	fmt.Println("secondmax value : \t ", min(list))
}
