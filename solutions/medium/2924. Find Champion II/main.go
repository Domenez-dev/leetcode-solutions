import (
	"container/list"
)

func findChampion(n int, edges [][]int) int {
	graph := make(map[int][]int)
	inDegree := make([]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		inDegree[v]++
	}

	var zeroInDegree []int
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			zeroInDegree = append(zeroInDegree, i)
		}
	}

	if len(zeroInDegree) == 0 {
		return -1
	}

	isChampion := func(node int) bool {
		visited := make(map[int]bool)
		queue := list.New()
		queue.PushBack(node)

		for queue.Len() > 0 {
			curr := queue.Remove(queue.Front()).(int)
			if !visited[curr] {
				visited[curr] = true
				for _, neighbor := range graph[curr] {
					if !visited[neighbor] {
						queue.PushBack(neighbor)
					}
				}
			}
		}

		return len(visited) == n
	}

	var candidates []int
	for _, node := range zeroInDegree {
		if isChampion(node) {
			candidates = append(candidates, node)
		}
	}

	if len(candidates) == 1 {
		return candidates[0]
	}
	return -1
}
