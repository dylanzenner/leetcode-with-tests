package unique

func Unique(n int) []int{
	res := make([]int, 0)
	count := 0

	i := 1
	for i < n {
		res = append(res, i)
		count += i
		i += 1
	}
	res = append(res, -count)
	return res
}