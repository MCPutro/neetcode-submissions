func twoSum(nums []int, target int) []int {
    mapTemp := make(map[int]int)

	for i, n := range nums{
		diff := target - n
		v, ok := mapTemp[diff]
		if ok {
			return []int{v, i}
		}
		mapTemp[n] = i
	}

	 return []int{}
}
