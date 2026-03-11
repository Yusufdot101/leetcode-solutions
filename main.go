package main

import "fmt"

func main() {
	nums := []int{4, 5, 6}
	target := 10

	fmt.Printf("is anagram: %v\n", twoSum(nums, target))
}
