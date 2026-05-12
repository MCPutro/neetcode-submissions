func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

// cara 1
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



// cara 2
	// sRunes, tRunes := []rune(s), []rune(t)

	// sort.Slice(sRunes, func(i, j int) bool {
    //     return sRunes[i] < sRunes[j]
    // })
	// sort.Slice(tRunes, func(i, j int) bool {
    //     return tRunes[i] < tRunes[j]
    // })

	// for i, _ := range sRunes{
	// 	if sRunes[i] != tRunes[i]{
	// 		return false
	// 	}
	// }

	return true

}
