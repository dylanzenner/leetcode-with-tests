package equaloccurrences

func EqualOccurrences(s string) bool {
	mapping := make(map[rune]int)

	for _, val := range s {
		_, ok := mapping[val]
		if !ok {
			mapping[val] = 1
		} else {
			mapping[val] += 1
		}
	}

	res := 0
	for _, v := range mapping {
		if res == 0 {
			res = v
		} else {
			if v != res {
				return false
			}
		}
	}
	return true
}