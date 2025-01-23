func maxArea(height []int) int {
	var min, steps, container, peak int

	for i := 0; i < len(height); i++ {
		if height[i] > peak {
			peak = height[i]
			for j := i; j < len(height); j++ {
				min = compare(height[i], height[j])
				steps = j - i
				container = max(container, min*steps)
			}
		}
	}
	return container
}

func compare(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
