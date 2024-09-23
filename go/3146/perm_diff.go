package permDiff

import "math"

func PermDiff(s string, t string) int {
	mappingOne := make(map[byte]int)
	mappingTwo := make(map[byte]int)

	res := 0
	pointer := 0

	for pointer < len(s) {
		mappingOne[s[pointer]] = pointer
		mappingTwo[t[pointer]] = pointer
		pointer += 1
	}

	pointer = 0

	for pointer < len(s) {
		res += int(math.Abs(float64(mappingOne[s[pointer]]) - float64(mappingTwo[s[pointer]])))
		pointer += 1
	}

	return res
}