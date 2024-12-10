package main

import "fmt"

func main() {
	n := 15

	fizzBuzz(n)
}

func fizzBuzz(num int) {
	for i:=1; i<=num ; i++{
		printFizzBuzz(i)
		if i < num {
			fmt.Print(", ")
		}
	}
}

func printFizzBuzz(n int) {
	switch {
	case n%5 == 0 && n%3 == 0:
		fmt.Print("FizzBuzz")
	case n%3==0 :
		fmt.Print("Fizz")
	case n%5==0 :
		fmt.Print("Buzz")
	default:
		fmt.Print(n)
	}
}