package main

import (
	"fmt"
)

func main() {
	// nums := []int{1, 2, 2, 3, 3, 3}
	// k := 2

	dummyInput := []string{"hello", "world"}
	s := Solution{}
	encoded := s.Encode(dummyInput)
	decoded := s.Decode(encoded)
	fmt.Printf("here: %v, %v\n", encoded, decoded)
}
