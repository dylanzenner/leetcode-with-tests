package matches

func NumMatches(n int) int {
	res := 0

	for n > 1 {
		if n % 2 == 0 {
			res += n / 2
			n = n / 2
		} else {
			res += (n - 1) / 2
			n = (n - 1) / 2 + 1
		}
	}

	return res
}