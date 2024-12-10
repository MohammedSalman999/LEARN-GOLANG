package main

import (
	"fmt"
	"sort"
)

func kthLargest(numbers []int, k int) int {
	sort.Ints(numbers)

	return numbers[len(numbers)-k]
}

func main() {
	numbers := []int{500, 450, 400, 350, 300, 250, 200, 150, 100, 50, 0}
	var k int
	fmt.Println("given list of int ", numbers)
	fmt.Println("Enter the number of your choice")
	fmt.Scanln(&k)

	switch {
	case k == 1:
		fmt.Printf("the %d'st largest element is %d", k, kthLargest(numbers, k))
	case k == 2:
		fmt.Printf("the %d'nd largest element is %d", k, kthLargest(numbers, k))
	case k == 3:
		fmt.Printf("the %d'rd' largest element is %d", k, kthLargest(numbers, k))
	default:
		fmt.Printf("the %d'th largest element is %d", k, kthLargest(numbers, k))
	}

}
