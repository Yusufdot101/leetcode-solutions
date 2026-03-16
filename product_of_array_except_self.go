package main

/**
Given an integer array nums, return an array output where output[i] is the product of all the elements of nums except nums[i].

Each product is guaranteed to fit in a 32-bit integer.

Follow-up: Could you solve it in O(n) time without using the division operation?

Example 1:
Input: nums = [1,2,4,6]
Output: [48,24,12,8]

Example 2:
Input: nums = [-1,0,1,2,3]
Output: [0,-6,0,0,0]

Constraints:
	2 <= nums.length <= 1000
	-20 <= nums[i] <= 20
**/

func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	product := 1
	for idx, val := range nums {
		product *= val
		prefix[idx] = product
	}

	suffix := make([]int, len(nums))
	product = 1
	for idx := len(nums) - 1; idx >= 0; idx-- {
		product *= nums[idx]
		suffix[idx] = product
	}

	result := make([]int, len(nums))
	for idx := range nums {
		left := 1
		if idx > 0 {
			left = prefix[idx-1]
		}

		right := 1
		if idx < len(nums)-1 {
			right = suffix[idx+1]
		}

		result[idx] = left * right
	}
	return result
}
