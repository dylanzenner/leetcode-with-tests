package peaks

func Peaks(mountain []int) []int {
	pointer := 1
	res := []int{}

	for pointer <= len(mountain) - 2 {
		if mountain[pointer] > mountain[pointer - 1] && mountain[pointer] > mountain[pointer + 1] {
			res = append(res, pointer)
		}
		pointer += 1
	}
	return res
}