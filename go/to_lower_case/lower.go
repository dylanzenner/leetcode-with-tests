package lower

func Lower(s string) string {
	res := []rune{}
	for _, val := range s {
		if val >= 65 && val <= 90 {
			val += 32
		}
		res = append(res, val)
	}
	return string(res)
}