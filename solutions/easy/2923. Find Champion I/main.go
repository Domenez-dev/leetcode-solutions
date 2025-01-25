import "fmt"

func findChampion(grid [][]int) int {
	l := len(grid)
	cand := 0
	for i := 1; i < l; i++ {
		if i != cand {
			if grid[cand][i] != 1 {
				cand = i
			}
		}
	}
	return cand
}
