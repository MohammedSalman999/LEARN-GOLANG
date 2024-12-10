package main

import "fmt"

func countEvan(list []int) (evan int) {
	if len(list) == 0 {
		return 0
	}
	for _, val := range list {
		if val%2 == 0 {
			evan++
		}
	}
	return evan
}

func main() {
	list := []int{2, 4, 5, 6, 8, 10}
	fmt.Println(countEvan(list))

}