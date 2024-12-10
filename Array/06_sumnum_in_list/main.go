package main

import "fmt"

func sumToList(li []int, num int) (int, int) {
	for i, val1 := range li {
		for j, val2 := range li {
			if i == j {
				continue
			}
			if val1+val2 == num {
				return i, j
			}
		}
	}

	return -1, -1
}

func main() {
	li := []int{1, 2, 3, 4, 5, 7}
	num := 5

	fmt.Println(sumToList(li, num))

}