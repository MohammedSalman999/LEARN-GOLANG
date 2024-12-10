package main

import (
	"fmt"
	"strings"
)
// simplest way accessing through range
func revStr(word string) string {
	result := ""

	for _, val := range word {
		result = string(val) + result
	}
	return result
}

// accesseing thorough loop 
func reverse(s string) (rev string) {
	for i := len(s)-1 ; i>=0; i-- {
		rev = rev + string(s[i])
	}
	return rev
}

func checkPalindrom(s string) bool {
	s = strings.ToLower(s)
	reversed := revStr(s)
	return reversed == s
}

func main() {
	s := "repaper"
	fmt.Println("word  \t \t ",s)
	fmt.Println("reverse  \t ",revStr(s))
	fmt.Println("reverse  \t ",reverse(s))
	fmt.Println("is Palindrom  \t ",checkPalindrom(s))
}