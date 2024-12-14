package main

import "fmt"

func main() {
	fmt.Println(ValidPalindrome("xdd"))
}
func ValidPalindrome(str string) bool {
	left := 0
	right := len(str) - 1

	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
