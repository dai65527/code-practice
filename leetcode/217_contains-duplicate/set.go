func containsDuplicate(nums []int) bool {
	dupmap := map[int]struct{}{}
	for i := range nums {
		if _, exists := dupmap[nums[i]]; exists {
			return true
		} else {
			dupmap[nums[i]] = struct{}{}
		}
	}
	return false
}
