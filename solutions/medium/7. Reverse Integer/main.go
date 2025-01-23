func reverse(x int) int {
	y := 0
	for x != 0 {
		y *= 10
		y += x % 10
		x /= 10
	}
	if y > 2147483647 || y < -2147483648 {
		return 0
	}
	return y
}
