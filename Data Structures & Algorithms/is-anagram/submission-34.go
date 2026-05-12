// import ("maps")
func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

	mapA := make(map[rune]int)
	mapB := make(map[rune]int)

	for i:=0; i<len(s);i++{
		mapA[rune(s[i])]++
		mapB[rune(t[i])]++
	}

	if len(mapA) != len(mapB){
		return false
	}

	for k, v := range mapB{
		if mapA[k] != v{
			return false
		}
	}
	// if !maps.Equal(mapA, mapB){
	// 	return false
	// }

// cara 1
	// mapS := make(map[rune]int)

	// for _, char := range s{
	// 	mapS[char]++
	// }

	// // mapT := make(map[rune]int)

	// // for _, char := range t{
	// // 	mapT[char]++
	// // }
	
	// // for k, v := range mapT{
	// // 	if mapS[k] != v{
	// // 		return false
	// // 	}
	// // }

	// for _, c := range t{
	// 	mapS[c]--
	// 	if mapS[c] < 0{
	// 		return false
	// 	}
	// }

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
