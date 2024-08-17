package diff

func Diff(s, t string) byte {
	var res rune
	mapping := make(map[rune]int)

	for _, val := range s {
		_, ok := mapping[val]
		if !ok {
			mapping[val] = 1
		} else {
			mapping[val] += 1
		}
	}

	for _, val := range t {
		if mapping[val] == 0 {
			res = val
		}
		mapping[val] -= 1
	}

	return byte(res)
}