// Date: March 7, 2020 (SAT)
// Link: https://leetcode.com/problems/two-sum/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mymap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		j, ok := mymap[target-nums[i]]
		if ok {
			result := []int{j, i}
			return result
		}
		mymap[nums[i]] = i
	}
	result := []int{-1, -1}
	return result
}

func main() {
	nums := []int{1, 3, 5, 7, 9}
	target := 4
	answer := twoSum(nums, target)
	fmt.Println(answer)
}
