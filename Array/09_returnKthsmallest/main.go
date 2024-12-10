package main

import (
	"fmt"
	"sort"
)

func kthSmallest(numbers []int, k int) int {
	sort.Ints(numbers)

	return numbers[k-1]
}

func main() {
	numbers := []int{500, 450, 400, 350, 300, 250, 200, 150, 100, 50, 0}
	k := 4

	fmt.Println("given list of int ", numbers)
	switch {
	case k == 1:
		fmt.Printf("the %d'st smallest element is %d", k, kthSmallest(numbers, k))
	case k == 2:
		fmt.Printf("the %d'nd smallest element is %d", k, kthSmallest(numbers, k))
	case k == 3:
		fmt.Printf("the %d'rd' smallest element is %d", k, kthSmallest(numbers, k))
	default:
		fmt.Printf("the %d'th smallest element is %d", k, kthSmallest(numbers, k))
	}

}
