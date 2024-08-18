package isacronym

func IsAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}

	pointer := 0
	for pointer < len(words) {
		if words[pointer][0] != s[pointer] {
			return false
		}
		pointer += 1
	}
	return true
}