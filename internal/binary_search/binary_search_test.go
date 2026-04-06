package binary_search_test

import (
	search "go_binary_search_example/internal/binary_search"
	generate "go_binary_search_example/internal/generate_nums"

	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

// HOW to run: go test

func TestBinarySearch_Positive(t *testing.T) {
	dataSet := generate.GenerateNums(10)
	slices.Sort(dataSet)

	testCases := []struct {
		name   string
		target int
		pos    int
		O      int
		err    error
	}{
		{
			name:   "target=0, pos=0, O(big)=2, err=<nil>",
			target: 0,
			pos:    0,
			O:      2,
			err:    nil,
		},
		{
			name:   "target=5, pos=5, O(big)=2, err=<nil>",
			target: 5,
			pos:    5,
			O:      2,
			err:    nil,
		},
		{
			name:   "target=1, pos=1, O(big)=2, err=<nil>",
			target: 1,
			pos:    1,
			O:      1,
			err:    nil,
		},
		{
			name:   "target=9, pos=9, O(big)=2, err=<nil>",
			target: 9,
			pos:    9,
			O:      3,
			err:    nil,
		},
		{
			name:   "target=11, pos=9, O(big)=2, err=<nil>",
			target: 11,
			pos:    -1,
			O:      4,
			err:    search.ErrNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pos, O, err := search.BinarySearch(dataSet, tc.target)

			require.Equal(t, tc.pos, pos)
			require.Equal(t, tc.O, O)
			require.Equal(t, tc.err, err)
		})
	}
}

// HOW to run: go test -bench=. -benchmem

func BenchmarkBinarySearch(b *testing.B) {
	nums := generate.GenerateNums(1_000_000)
	slices.Sort(nums) // Binary search works when the slice is sorted
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, _ = search.BinarySearch(nums, i)
	}
}
