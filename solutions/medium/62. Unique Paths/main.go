func uniquePaths(m int, n int) int {
	// move right m--
	// move down n--
	moves := 0
	var cache map[[2]int]int
	cache = make(map[[2]int]int)

	total_ways := moving(m, n, moves, cache)
	return total_ways
}

func moving(m, n, moves int, cache map[[2]int]int) int {
	key := [2]int{m, n}
	if n == 1 && m == 1 {
		return 1
	}

	if m == 0 || n == 0 {
		return 0
	}

	if cache[key] != 0 {
		return cache[key]
	}

	// Right cell + Down cell
	moves = moving(m-1, n, moves, cache) + moving(m, n-1, moves, cache)

	// store value to cache
	inverted_key := [2]int{n, m} // for inverted key since moving[m,n] == moving[n,m]
	cache[key] = moves
	cache[inverted_key] = moves
	return moves
}
