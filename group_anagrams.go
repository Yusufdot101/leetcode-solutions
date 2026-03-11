package main

import (
	"slices"
	"strings"
)

/**
	Given an array of strings strs, group all anagrams together into sublists. You may return the output in any order.
	An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

	Input: strs = ["act","pots","tops","cat","stop","hat"]
	Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]

	Input: strs = ["x"]
	Output: [["x"]]

	Input: strs = [""]
	Output: [[""]]


	Constraints:
		1 <= strs.length <= 1000.
		0 <= strs[i].length <= 100
		strs[i] is made up of lowercase English letters.


	Input: strs = ["abc", "cab", bac]
**/

func groupAnagrams(strs []string) [][]string {
	ogStrs := make([]string, len(strs))
	copy(ogStrs, strs)
	for idx, str := range strs {
		arr := strings.Split(str, "")
		slices.Sort(arr)
		str := strings.Join(arr, "")
		strs[idx] = str
	}

	dict := make(map[string][]int)
	for idx, str := range strs {
		dict[str] = append(dict[str], idx)
	}

	solution := [][]string{}
	subSolution := []string{}
	for _, v := range dict {
		for _, idx := range v {
			subSolution = append(subSolution, ogStrs[idx])
		}
		solution = append(solution, subSolution)
		subSolution = []string{}
	}

	return solution
}
