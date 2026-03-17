package easy

func TwoSum(nums []int, target int) []int {
	solution := []int{}
	dict := make(map[int]int)
	for idx, num := range nums {
		compliment := target - num
		if v, exists := dict[compliment]; exists {
			solution = append(solution, v, idx)
			break
		}
		dict[num] = idx
	}
	return solution
}
