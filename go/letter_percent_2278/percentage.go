package percentage

func Percentage(s string, letter byte) int {
	count := 0
	for _, val := range s {
		if byte(val) == letter {
			count += 1
		}
	}

	if count == 0{
		return 0
	} else {
		return int( (float64(count) / float64(len(s)) * 100))
	}
}