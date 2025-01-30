package main

func SearchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1 // for left and right

	for l <= r {
		i := l + (r-l)/2 // i for middle index
		if nums[i] == target {
			return i
		} else if nums[i] < target {
			l = i + 1
		} else {
			r = i - 1
		}
	}

	return l
}
