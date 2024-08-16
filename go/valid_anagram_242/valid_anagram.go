package validAnagram

func ValidAnagram(s, t string) bool {
	mapping := make(map[byte]int)

	pointer := 0

	if len(s) != len(t){
		return false
	}

	for pointer < len(s){
		_, ok := mapping[s[pointer]]
		if !ok {
			mapping[s[pointer]] = 1
		} else {
			mapping[s[pointer]] += 1
		}
		pointer += 1
	}

	pointer = 0
	for pointer < len(t){
		_, ok := mapping[t[pointer]]
		if !ok {
			return false
		} else {
			mapping[t[pointer]] -= 1
		}
		pointer += 1
	}

	pointer = 0
	for pointer < len(s){
		if mapping[s[pointer]] != 0{
			return false
		}
		pointer += 1
	}
	return true
}