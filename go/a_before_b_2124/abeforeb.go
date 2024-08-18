package abeforeb

func AbeforeB(s string) bool {
	left := 0
	right := len(s) - 1

	a_pos := -1
	b_pos := len(s) + 1

	for left < len(s) && right > -1 {
		if s[right] == 'a' && right > a_pos {
			a_pos = right
		}

		if s[left] == 'b' && left < b_pos {
			b_pos = left
		}

		right -= 1
		left += 1
	}

	return a_pos < b_pos
}