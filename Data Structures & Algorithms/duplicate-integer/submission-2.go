func hasDuplicate(nums []int) bool {
    mapSet := make(map[int]bool)

	for _, v := range nums{
		if mapSet[v]{
			return true
		}
		mapSet[v]=true
	}

	return false
}
