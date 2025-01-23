func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	x, y := len(nums1), len(nums2)
	low, high := 0, x

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (x+y+1)/2 - partitionX

		maxX := math.MinInt32
		if partitionX > 0 {
			maxX = nums1[partitionX-1]
		}

		maxY := math.MinInt32
		if partitionY > 0 {
			maxY = nums2[partitionY-1]
		}

		minX := math.MaxInt32
		if partitionX < x {
			minX = nums1[partitionX]
		}

		minY := math.MaxInt32
		if partitionY < y {
			minY = nums2[partitionY]
		}

		if maxX <= minY && maxY <= minX {
			if (x+y)%2 == 0 {
				return float64(max(maxX, maxY)+min(minX, minY)) / 2.0
			} else {
				return float64(max(maxX, maxY))
			}
		} else if maxX > minY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}

	panic("Input arrays are not sorted")
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
