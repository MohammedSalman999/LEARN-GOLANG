package main

import "fmt"

func freq(arr []int, n int) int {
	result := make(map[int]int)
	for _, num := range arr {
		result[num]++
	}
	return result[n]
}

func main() {
	arr := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 7, 8, 8, 8, 8, 9, 9, 9, 0}
	n := 1
	fmt.Println("the array given is ", arr)
	fmt.Printf("%d came %d times ", n, freq(arr, n))
}
