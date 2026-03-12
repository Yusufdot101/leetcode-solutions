package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 3, 3}
	k := 2

	fmt.Printf("output: %v\n", topKFrequent(nums, k))
}
