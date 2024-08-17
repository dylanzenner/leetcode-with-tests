package subseq

func SubSeq(s, t string) bool {
	left_pointer := 0
	right_pointer :=0

	if len(s) == 0{
		return true
	}

	for right_pointer < len(t) && left_pointer < len(s) {
		if s[left_pointer] == t[right_pointer] {
			left_pointer += 1
		}
		right_pointer +=1
	}

	if left_pointer == len(s) {
		return true
	} else {
		return false
	}
}