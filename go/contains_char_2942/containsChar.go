package containschar

func ContainsChar(words []string, x byte) []int {
	pointer := 0
	res := []int{}

	for pointer < len(words) {
		for _, val := range words[pointer] {
			if byte(val) == x {
				res = append(res, pointer)
				break
			}
		}
		pointer += 1
	}
	return res
}