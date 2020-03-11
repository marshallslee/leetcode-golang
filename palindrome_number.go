// Date: March 8, 2020 (SUN)
// Link: https://leetcode.com/problems/palindrome-number/

package main

import "fmt"

func isPalindrome(x int) bool {
	var remainder, temp int
	reverse := 0

	temp = x

	for {
		remainder = x % 10
		reverse = reverse*10 + remainder
		x /= 10

		if x == 0 {
			break
		}
	}

	return temp == reverse
}

func main() {
	num := 10001
	palindrome := isPalindrome(num)
	fmt.Println(palindrome)
}
