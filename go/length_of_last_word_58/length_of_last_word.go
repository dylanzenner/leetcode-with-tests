package lengthOfLastWord

import "strings"

func LengthOfLastWord(s string) int {
	count := 0
	trimS := strings.Trim(s, " ")
	pointer := len(trimS) - 1

	for true {
		if pointer < 0 {
			return count
		}

		if string(trimS[pointer]) == " " {
			return count
		} else {
			count += 1
			pointer -= 1
		}
	}

	return count
}