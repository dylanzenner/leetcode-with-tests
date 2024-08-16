package intersect

import "slices"

func Intersect(num1, num2 []int) []int {
	res := []int{}
	for _, val := range(num1){
		if slices.Contains(num2, val) && !slices.Contains(res, val) {
			res = append(res, val)
		}
	}
	return res
}