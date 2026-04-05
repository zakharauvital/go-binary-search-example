package main

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
)

func generateNums(n int) []int {
	items := make([]int, 0, n)

	for i := range n {
		items = append(items, i)
	}

	return items
}

func binarySearch(nums []int, target int) (int, int, error) {
	O := 0
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2 // find a middle of a slice
		guessed := nums[mid]

		switch true {
		case guessed == target:
			return mid, O, nil
		case guessed > target: // too big
			right = mid - 1
		case guessed < target: // too small
			left = mid + 1
		}

		O++
	}

	return -1, O, errors.New("not found")
}

// HOW to run: "go run main.go"

func main() {
	nums := generateNums(1_000_000)
	slices.Sort(nums) // Binary search works when the slice is sorted
	pos, O, err := binarySearch(nums, rand.Intn(len(nums)))
	fmt.Printf("pos=%d, O(log n)=%d, err=%v\n", pos, O, err) // "O-Big" describes the algorithm’s speed
}
