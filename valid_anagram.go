package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sDict := make(map[rune]int)
	tDict := make(map[rune]int)

	for _, char := range s {
		sDict[char] += 1
	}

	for _, char := range t {
		tDict[char]++
	}

	for k, v := range sDict {
		if tDict[k] != v {
			return false
		}
	}
	return true
}
