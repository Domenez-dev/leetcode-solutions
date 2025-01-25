func diagonalSum(mat [][]int) int {
	sum := 0
	l := len(mat)
	for i := 0; i < l; i++ {
		sum = sum + mat[i][i] + mat[i][l-i-1]
	}
	if l%2 != 0 {
		mid := l / 2
		sum -= mat[mid][mid]
	}
	return sum
}
