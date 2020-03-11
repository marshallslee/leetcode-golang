// Date: March 8, 2020 (SUN)
// Link: https://leetcode.com/problems/palindrome-number/

package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func main() {
	num := 10001
	palindrome := isPalindrome(num)
	fmt.Println(palindrome)
}
