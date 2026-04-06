package binary_search

import (
	"errors"
)

var ErrNotFound = errors.New("not found")

func BinarySearch(nums []int, target int) (int, int, error) {
	O := 0
	mid := 0
	left, right := 0, len(nums)-1

	for left <= right {
		mid = (left + right) / 2 // find a middle of a slice
		guessed := nums[mid]

		switch true {
		case guessed == target: // found
			return mid, O, nil
		case guessed > target: // too big
			right = mid - 1
		default: // too small
			left = mid + 1
		}

		O++
	}

	return -1, O, ErrNotFound
}
