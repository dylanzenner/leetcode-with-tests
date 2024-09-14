package sumMultiple

func SumMultiple(n int) int {
	count := 0
	for i := range n + 1 {
		if i % 3 == 0 || i % 5 == 0 || i % 7 == 0 {
			count += i
		}
	}
	return count
}