package main

import (
	"fmt"
	"strconv"
)

func main() {
	flag := IsPalindrome(10101)
	fmt.Println(flag)
}

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	s := strconv.Itoa(x)
	len := len(s)
	h1 := len / 2

	for i := 0; i <= h1; i++ {
		if s[i:i+1] != s[len-i-1:len-i-1+1] {
			return false
		}
	}
	return true
}
