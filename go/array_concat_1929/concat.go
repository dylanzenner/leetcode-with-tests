package concat

func Concat(nums []int) []int {
	pointer := 0
	count := 2
	res := []int{}

	for count > 0{
		if pointer == len(nums) {
			pointer = 0
			count -= 1
			if count == 0 {
				break
			}
		}
		res = append(res, nums[pointer])
		pointer += 1
	}
	return res
}