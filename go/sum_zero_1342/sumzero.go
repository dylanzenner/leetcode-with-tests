package sumzero

func SumZero(n int) int{
	count := 0 
	for n != 0 {
		if n % 2 == 0 {
			n /= 2
			count += 1
		} else {
			n -= 1
			count += 1
		}
	}
	return count
}