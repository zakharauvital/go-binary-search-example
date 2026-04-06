package generatenums

func GenerateNums(n int) []int {
	items := make([]int, 0, n)

	for i := range n {
		items = append(items, i)
	}

	return items
}
