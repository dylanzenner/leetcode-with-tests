package jewels

func FindJewels(jewels, stones string) int {
	mapping := make(map[rune]bool)
	count := 0

	for _, val := range jewels {
		_, ok := mapping[val]
		if !ok {
			mapping[val] = true
		}
	}

	for _, val := range stones {
		_, ok := mapping[val]
		if ok {
			count += 1
		}
	}

	return count
}