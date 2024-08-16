package reverseString

func ReverseString(s []byte) []byte {
	left_pointer := 0
	right_pointer := len(s) - 1

	for left_pointer < right_pointer {
		s[left_pointer], s[right_pointer] = s[right_pointer], s[left_pointer]
		left_pointer += 1
		right_pointer -= 1
	}

	return s
}