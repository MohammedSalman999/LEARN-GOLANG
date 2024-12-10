package main

import "fmt"

func twosum(slice []int, target int) (int, int, bool) {
	m := make(map[int]int)
	for i, num := range slice {
		remender := target - num
		if j, ok := m[remender]; ok {
			return slice[j], num, true
		}
		m[num] = i
		fmt.Println(m)
	}
	return 0, 0, false
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("given slice ", slice)
	target := 8
	n1, n2, ok := twosum(slice, target)
	fmt.Println("two sum num is ", n1, n2, ok)
}
