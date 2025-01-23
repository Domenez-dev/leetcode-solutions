import (
	"fmt"
	"math"
)

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i+1 == j {
				graph[i][j] = 1
			} else {
				graph[i][j] = math.MaxInt32
			}
		}
	}

	dijkstra := func() int {
		dist := make([]int, n)
		visited := make([]bool, n)

		for i := range dist {
			dist[i] = math.MaxInt32
		}
		dist[0] = 0

		for count := 0; count < n; count++ {

			minDist := math.MaxInt32
			u := -1
			for i := 0; i < n; i++ {
				if !visited[i] && dist[i] < minDist {
					minDist = dist[i]
					u = i
				}
			}

			if u == -1 {
				break
			}

			visited[u] = true

			for v := 0; v < n; v++ {
				if graph[u][v] != math.MaxInt32 && !visited[v] {
					newDist := dist[u] + graph[u][v]
					if newDist < dist[v] {
						dist[v] = newDist
					}
				}
			}
		}
		return dist[n-1]
	}

	answer := make([]int, 0, len(queries))
	for _, query := range queries {
		u, v := query[0], query[1]
		graph[u][v] = 1
		answer = append(answer, dijkstra())
	}

	return answer
}
