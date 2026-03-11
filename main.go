package main

import "fmt"

func main() {
	s := "racecar"
	t := "carrace"

	fmt.Printf("is anagram: %v\n", isAnagram(s, t))
}
