func twoSum(nums []int, target int) []int {
	mp := map[int]int{}
	for i, v := range nums {
		if j, ok := mp[target-v]; ok {
			return []int{j, i}
		}
		mp[v] = i
	}
	return nil
}
