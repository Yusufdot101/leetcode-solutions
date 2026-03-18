package easy

import (
	"slices"
)

/**
You are given a string s consisting of the following characters: '(', ')', '{', '}', '[' and ']'.

The input string s is valid if and only if:
    Every open bracket is closed by the same type of close bracket.
    Open brackets are closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.
Return true if s is a valid string, and false otherwise.

Example 1:
Input: s = "[]"
Output: true

Example 2:
Input: s = "([{}])"
Output: true

Example 3:
Input: s = "[(])"
Output: false
Explanation: The brackets are not closed in the correct order.

Constraints:
    1 <= s.length <= 1000
**/

func IsValidParenthesis(s string) bool {
	if len(s) == 1 {
		return false
	} else if len(s) == 0 {
		return true
	}

	stack := []rune{}
	closing := []rune{']', '}', ')'}
	hashMap := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}

	for _, par := range s {
		if slices.Contains(closing, par) {
			if len(stack) == 0 || stack[len(stack)-1] != hashMap[par] {
				return false
			}
			// pop the last
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, par)
	}

	return len(stack) == 0
}
