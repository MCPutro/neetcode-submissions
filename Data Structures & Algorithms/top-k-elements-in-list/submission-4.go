func topKFrequent(nums []int, k int) []int {
	mapSet := make(map[int]int)

	for _ , n := range nums {
		mapSet[n]++
	}

	mapGroup := map[int][]int{}

	for key, freq := range mapSet {
		mapGroup[freq] = append(mapGroup[freq], key)
	}

	var keys []int
	for key := range mapGroup {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	var resp []int

	for i := len(keys); i > 0; i-- {
		resp = append(resp, mapGroup[keys[i-1]]...)
	}

	fmt.Println(resp)
	return resp[:k]

}
