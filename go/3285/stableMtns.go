package stableMtns

func StableMtns(height []int, threshold int) []int {
	res := []int{}
	left, right := 0, 1

	for right < len(height) {
		if height[left] > threshold {
			res = append(res, right)
		}
		left += 1
		right += 1
	}
	return res
}