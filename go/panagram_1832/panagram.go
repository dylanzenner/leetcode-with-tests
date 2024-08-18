package panagram

func Panagram(s string) bool {
	if len(s) < 26 {
		return false
	}

	mapping := make(map[rune]int)

	for _, val := range s {
		_, ok := mapping[val]
		if !ok {
			mapping[val] = 1
		}
	}
	return len(mapping) == 26
}