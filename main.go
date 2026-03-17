package main

import (
	"fmt"

	"github.com/Yusufdot101/leetcode-solutions/intermediate"
)

func main() {
	nums := []int{2, 20, 4, 10, 3, 4, 5}
	output := intermediate.LongestConsecutive(nums)
	fmt.Println("output: ", output)

	nums2 := []int{0, 3, 2, 5, 4, 6, 1, 1}
	output2 := intermediate.LongestConsecutive(nums2)
	fmt.Println("output: ", output2)
}
