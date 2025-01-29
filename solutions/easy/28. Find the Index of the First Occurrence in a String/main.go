func strStr(haystack string, needle string) int {
	if !(len(haystack) < len(needle)) {
		for i := 0; i <= len(haystack)-len(needle); i++ {
			for j := 0; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					break
				}
				if (j == len(needle)-1) && haystack[i+j] == needle[j] {
					return i
				}
			}
		}
	}
	return -1
}
