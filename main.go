package main

import (
	"fmt"
	search "go_binary_search_example/internal/binarysearch"
	generate "go_binary_search_example/internal/generatenums"
	"math/rand"
	"slices"
)

// HOW to run: "go run main.go"

func main() {
	nums := generate.GenerateNums(1_000_000)
	slices.Sort(nums) // Binary search works when the slice is sorted

	pos, O, err := search.BinarySearch(nums, rand.Intn(len(nums)))

	fmt.Printf("pos=%d, O(log n)=%d, err=%v\n", pos, O, err) // "O-Big" describes the algorithm’s speed
}
