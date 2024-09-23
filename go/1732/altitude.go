package altitude

import "math"

func Altitude(nums []int) int {
	res := math.Inf(-1)
	val := 0.0

	for _, v := range nums {
		val += float64(v)
		if val > res {
			res = val
		}
	}

	if res < 0 {
		return 0
	} else {
		return int(res)
	}
}