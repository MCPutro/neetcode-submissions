func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

	mapS := make(map[rune]int)

	for _, char := range s{
		mapS[char]++
	}

	mapT := make(map[rune]int)

	for _, char := range t{
		mapT[char]++
	}
	
	for k, v := range mapT{
		if mapS[k] != v{
			return false
		}
	}

	return true

}
