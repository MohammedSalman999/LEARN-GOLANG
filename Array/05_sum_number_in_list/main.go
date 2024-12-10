package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5}

	fmt.Println(list)
	fmt.Println(sum(list))
	fmt.Println(Sum(list))
}

// using range 

func sum(list []int) (num int) {
	for _, val := range list {
		num = num + val
	}

	return num
}






// recursion

func Sum(list []int) int {
	if len(list) == 0 {
		return 0
	}

	return list[0] + Sum(list[1:])
}