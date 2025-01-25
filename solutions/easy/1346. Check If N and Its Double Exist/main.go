func checkIfExist(arr []int) bool {
	if len(arr) == 0 || len(arr) == 1 {
		return false
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if (arr[i] == arr[j]*2) || (arr[j] == arr[i]*2) {
				return true
			}
		}
	}
	return false
}
