package finalval

func FinalVal(operations []string) int {
	res := 0

	for _, op := range operations {
		if op == "++X" || op == "X++" {
			res += 1
		} else {
			res -= 1
		}
	}
	return res
}