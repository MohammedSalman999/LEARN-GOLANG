package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println(revArr(arr))
}

func revArr(arr []int) []int {
	left , right := 0 , len(arr)-1

	for left < right {
		arr[left], arr[right] = arr[right] , arr[left]
		left ++
		right --
	}

	return arr
}