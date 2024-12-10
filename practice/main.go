// package main

// import "fmt"

// func secondMax(numbers []int) int {
// 	max := numbers[0]
// 	var second_max int

// 	for _, num := range numbers {
// 		if num > max {
// 			second_max = max
// 			max = num
// 		} else if num > second_max && num != max {
// 			second_max = num
// 		}
// 	}
// 	return second_max
// }

// func main() {

// 	numbers := []int{5, 1, 2, 3, 4}
// 	fmt.Println("numbers \t : ", numbers)
// 	fmt.Println("second max : ", secondMax(numbers))
// }

// package main

// import "fmt"

// // for sum all the elements in array
// // hume loop chalakar har element ko jodna hoga
// func sumNum(numbers []int) int {
// 	result := 0

// 	for _, num := range numbers {
// 		result = result + num
// 	}
// 	return result
// }

// func main() {
// 	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	fmt.Println(sumNum(numbers))
// }

// find the secondmax value in an array

// package main

// import "fmt"

// func arrayValue(numbers []int) int {
// 	max := numbers[0]
// 	secondMax := 0
// 	for _, num := range numbers {
// 		if num > max {
// 			secondMax = max
// 			max = num
// 		} else if num > secondMax && num != max {
// 			secondMax = num
// 		}
// 	}
// 	return secondMax
// }

// func main() {
// 	numbers := []int{1, 2, 3}
// 	fmt.Println(arrayValue(numbers))

// }

// package main

// import "fmt"

// //write a function to find the best day to buy the share
// // and the best day to sell it

// type money int

// func (m money) String() string {
// 	return fmt.Sprintf("$ %d", m)
// }
// func findDay(slice []int) (money, money, money, money, money) {
// 	buyDay := 0
// 	buyPrice := slice[0]
// 	sellDay := 0
// 	sellPrice := 0

// 	for day, price := range slice {
// 		if price < buyPrice {
// 			buyPrice = price
// 			buyDay = day
// 		} else if profit := price - buyPrice; profit > sellPrice {
// 			sellPrice = profit
// 			sellDay = day
// 		}
// 	}
// 	return money(buyDay), money(buyPrice), money(sellDay), money(slice[sellDay]), money(sellPrice)
// }

//	func main() {
//		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//		fmt.Printf("given list : \t %v\n", slice)
//		buyDay, buyPrice, sellDay, sellPrice, profit := findDay(slice)
//		fmt.Printf("final decision : \t\n\t\t buy day : \t %d \t buy price : \t %s \n\t\t sellday : \t %d \t sellprice :\t %s \n\t\t profit : \t %s", buyDay+1, buyPrice, sellDay+1, sellPrice, profit)
//	}
// package main

// import "fmt"

// type money int

// func (m money) String() string {
// 	return fmt.Sprintf("$ %d ", m)
// }

// func find(slice []int) (money, money, money, money, money) {
// 	bday := 0
// 	bprice := slice[0]
// 	sday := 0
// 	maxPrice := 0

// 	for day, price := range slice {
// 		if price < bprice {
// 			bprice = price
// 			bday = day
// 		} else if profit := price - bprice; profit > maxPrice {
// 			maxPrice = profit
// 			sday = day
// 		}
// 	}
// 	return money(bday), money(bprice), money(sday), money(slice[sday]), money(maxPrice)
// }

// func main() {
// 	stocks := []int{1, 3, 5, 4, 9}
// 	fmt.Println("give stocks are : \n\t\t ", stocks)
// 	bd, bprice, sd, sprice, profit := find(stocks)
// 	fmt.Printf("final decision : \n\t\t buy day %d at price %s \n\t\t sell day %d at price %s \n\t\t profit earn %d", bd+1, bprice, sd+1, sprice, profit)
// }

// package main

// import "fmt"

// // write a program to find the frequncy of the number
// func freq(arr []int, n int) int {
// 	m := make(map[int]int)

// 	for _, num := range arr {
// 		if num == n {
// 			m[num]++
// 		}
// 	}
// 	return m[n]
// }

//	func main() {
//		arr := []int{11, 11, 11, 11, 11, 12, 12, 13, 13, 14, 14, 15, 15, 16, 17, 18, 16, 19, 20}
//		n := 16
//		fmt.Println("given arr :", arr)
//		fmt.Printf("%d repeated %d times", n, freq(arr, n))
//	}
// package main

// import "fmt"

// // write a program to find the second last element of an array
// func find(arr []int) int {
// 	max := arr[0]
// 	smax := 0
// 	if len(arr) < 2 {
// 		return 0
// 	}

// 	switch {
// 	case arr[1] > arr[0]:
// 		max = arr[1]
// 		smax = arr[0]
// 	default:
// 		max = arr[0]
// 		smax = arr[1]
// 	}

// 	for _, num := range arr {
// 		if num > max {
// 			smax = max
// 			max = num
// 		} else if num > smax && num != max {
// 			smax = num
// 		}
// 	}
// 	return smax
// }

// func main() {
// 	arr := []int{1, 2}
// 	fmt.Println("give array", arr)
// 	fmt.Printf("%d is the second max number", find(arr))
// }

// write a program to reverse a string

// package main

// func rev(s string) string {
// 	rev := "" // O(1)
// 	// here we are ranging over the runes in slice
// 	// strings underslying type is
// 	// ye read only []bytes , and runes typically works like that
// 	for _, r := range s {
// 		rev = string(r) + rev
// 	}
// 	return rev
// }

// func revs(s string) string {
// 	runes := []rune(s)
// 	n := len(runes) //5-0-1 == 4
// 	for i := 0; i < n/2; i++ {
// 		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
// 	}
// 	return string(runes)
// }

// package main

// import "fmt"

// // write a progarm to find the common fizzbuzz problem

// func fizzbuzz(num int) {
// 	for i := 1; i <= num; i++ {
// 		printFizzBuzz(i)
// 		if i != num {
// 			fmt.Print(", ")
// 		}
// 	}

// }

// func printFizzBuzz(n int) {
// 	switch {
// 	case n%5 == 0 && n%3 == 0:
// 		fmt.Print("fizzBuzz")
// 	case n%3 == 0:
// 		fmt.Print("fizz")
// 	case n%5 == 0:
// 		fmt.Print("Buzz")
// 	default:
// 		fmt.Print(n)
// 	}
// }

// func main() {
// 	n := 15
// 	fizzbuzz(n)
// }

// package main

// import "fmt"

// func fabinacci(n int) []int {
// 	fab := make([]int, n)
// 	if len(fab) < 2 {
// 		return []int{}
// 	}

// 	fab[0] = 0
// 	fab[1] = 1
// 	for i := 2; i < len(fab); i++ {
// 		fab[i] = fab[i-1] + fab[i-2]
// 	}

// 	return fab
// }

// func main() {
// 	numebr := 2
// 	fmt.Println(fabinacci(numebr))
// }

// package main

// import "fmt"

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	n := 12
// 	numebrs := twosum(arr, n)
// 	i := 1
// 	for k, v := range numebrs {
// 		fmt.Printf("%d pair is %d , %d \n", i, k, v)
// 		i++
// 	}
// }

// // function two find the pairs for the given value

// func twosum(a []int, n int) map[int]int {
// 	m := make(map[int]int)
// 	ret := make(map[int]int)
// 	for i, num := range a {
// 		compl := n - num
// 		if j, ok := m[compl]; ok {
// 			ret[a[j]] = num
// 		}
// 		m[num] = i
// 	}
// 	return ret
// }

// package main

// import "fmt"

// // write a program to reverse an array

// func main() {
// 	ar := [5]int{1, 2, 3, 4, 5}
// 	fmt.Println(ar)
// 	fmt.Println(revArr(ar))
// }

// func revArr(ar [5]int) [5]int {
// 	f := 0
// 	l := len(ar) - 1

// 	for f < l {
// 		ar[f], ar[l] = ar[l], ar[f]
// 		f++
// 		l--
// 	}
// 	return ar
// }

// package main

// import "fmt"

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
// 	target := 14

// 	fmt.Println("given array is given below", arr)
// 	result := pair(arr, target)
// 	fmt.Println("the pairs are as follows", result)
// }

// func pair(ar []int, t int) [][2]int {
// 	m := make(map[int]int)
// 	var ret [][2]int
// 	for i, num := range ar {
// 		complimentry := t - num
// 		if j, ok := m[complimentry]; ok {
// 			ret = append(ret, [2]int{ar[j], num})
// 		}
// 		m[num] = i
// 	}
// 	return ret
// }

// write a function to find two nums that sum to a given number

// package main

// import "fmt"

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	target := 15
// 	fmt.Println("the pairs are : ")
// 	for key, value := range twosum(arr, target) {
// 		fmt.Printf("% d + %d = %d\n", key, value, key+value)
// 	}
// }

// func twosum(numbers []int, target int) map[int]int {
// 	m := make(map[int]int)
// 	result := make(map[int]int)

// 	for i, num := range numbers {
// 		cmpl := target - num
// 		if j, ok := m[cmpl]; ok {
// 			result[numbers[j]] = num
// 		}
// 		m[num] = i
// 	}

//		return result
//	}
// package main

// import "fmt"

// type money int

// // lets make custom type
// func (m money) String() string {
// 	return fmt.Sprintf("$%d", m)
// }

// // target : write a fucniton to find the best day to buy and share share
// // soo we need prices and the days of the shares for the but day nad selling day
// //function deintaion and logics :
// func getday(array []int) (int, int, int, int, int) {
// 	var (
// 		bprice           = array[0]
// 		bd, sprice, sday int
// 		maxProfit        int
// 	)
// 	for day, price := range array {
// 		if price < bprice {
// 			bprice = price
// 			bd = day
// 		} else if profit := price - bprice; profit > maxProfit {
// 			maxProfit = profit
// 			sprice = price
// 			sday = day
// 		}
// 	}
// 	return bd, bprice, sday, sprice, maxProfit
// }

// func main() {
// 	array := []int{100, 20, 30, 40, 10, 50, 60, 70, 80, 90}
// 	bd, bprice, sd, sp, mp := getday(array)
// 	fmt.Printf("the best day to buy the share is day %d at price %s \n", bd+1, money(bprice))
// 	fmt.Printf("and the best day to sell the share is day %d at price %s \n", sd+1, money(sp))
// 	fmt.Printf("and the profit is %s", money(mp))
// }

// write a program to find the best day to buy the share

// package main

// import "fmt"

// func main() {
// 	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
// 	bday, bprice, sellDay, sellPrice := findDay(arr)
// 	fmt.Printf("the best day to buy the stock is %d , at price of rupees %s \n", bday+1, money(bprice))
// 	fmt.Printf("and the best day to sell the stock is %d at the price of %s \n", sellDay+1, money(sellPrice))

// }

// type money int

// func (m money) String() string {
// 	return fmt.Sprintf("$%d", m)
// }

// func findDay(arr []int) (money, money, money, money) {
// 	bPrice := arr[0]
// 	var (
// 		bday      int
// 		sellDay   int
// 		sellprice int
// 	)

// 	for day, price := range arr {
// 		if price < bPrice {
// 			bPrice = price
// 			bday = day
// 		} else if profit := price - bPrice; profit > int(sellprice) {
// 			sellprice = profit
// 			sellDay = day
// 		}
// 	}
// 	return money(bday), money(bPrice), money(sellDay), money(sellprice)
// }

// package main

// import "fmt"

// // write a program to find the frequency of the element
// func freq(arr []int, target int) int {
// 	m := make(map[int]int)
// 	for _, num := range arr {
// 		m[num]++
// 	}
// 	return m[target]
// }

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 56, 7, 8, 1, 1, 1, 2, 2, 2}
// 	target := 2
// 	fmt.Print("the number ", target, " came times ", freq(arr, target))
// }

// package main

// import "fmt"

// // write a program to find the reverse of the string

// func reverse(word string) (string, bool) {
// 	rev := ""
// 	for _, s := range word {
// 		rev = string(s) + rev
// 	}
// 	return rev, word == rev
// 	//  check wheater its a palindrom or not

// }

// func main() {
// 	word := "hello "
// 	rev, isPal := reverse(word)
// 	fmt.Println("the reverse of the given string is ", rev)
// 	fmt.Println("the given string is a palindrom or not ", isPal)
// }

// write a program to find the sum of the given array

// package main

// import "fmt"

// func sum(arr []int) int {
// 	// given array is
// 	result := 0
// 	for _, num := range arr {
// 		result = result + num
// 	}
// 	return result
// }
// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	fmt.Printf("the sum of the given array is %d ", sum(arr))
// }

// write a program to find the second maximum value of an array
package main

import "fmt"

func secondMax(arr []int) int {
	max := 0
	secondmax := 0
	for _, num := range arr {
		if num > max {
			secondmax = max
			max = num
		} else if num > secondmax && num != max {
			secondmax = num
		}
	}
	return secondmax
}

func main() {
	given := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println("the given array is ")
	fmt.Println("the second maximum vlaue of an aray is ", secondMax(given))
}
