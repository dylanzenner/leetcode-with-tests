package containsDuplicate

func ContainsDuplicate(nums []int) bool {
	lookUp := make(map[int]bool)

	for i := range len(nums){
		lookUp[nums[i]] = true

	}
	return len(lookUp) != len(nums)
}