func isPalindrome(s string) bool {
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	clean := reg.ReplaceAllString(s, "")
	clean = strings.ToLower(clean)
	fmt.Println(">>",clean)
	length := len(clean)
	for i, j := 0, length-1; i < length/2; i, j = i+1, j-1 {
		if clean[i] != clean[j]{
			return false
		}
	}
	return true
}
