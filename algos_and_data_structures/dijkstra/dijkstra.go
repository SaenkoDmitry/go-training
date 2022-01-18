package dijkstra

import "container/heap"

type Graph map[string]map[string]int

func (g *Graph) GetDijkstraPath(start, target string) (path []string, cost int, err error) {
	h := NewMinHeap()
	heap.Push(h, &Node{
		key:  start,
		cost: 0,
	})
	visited := make(map[string]bool)
	previous := make(map[string]string)
	costs := make(map[string]int)

	for h.Len() > 0 {
		top := heap.Pop(h).(*Node)

		// 1) if found
		if top.key == target {
			cost = top.cost
			key := top.key
			for key != start {
				path = append(path, key)
				key = previous[key]
			}
			break
		}

		// 2) mark visited
		visited[top.key] = true

		// 3) visit all neighboring nodes
		for nextKey, nextCost := range (*g)[top.key] {
			if visited[nextKey] {
				continue
			}

			if minCost, ok := costs[nextKey]; !ok || minCost > top.cost + nextCost {
				heap.Push(h, &Node{key: nextKey, cost: top.cost + nextCost})
				previous[nextKey] = top.key
			}
		}
	}

	path = append(path, start)

	for i, j := 0, len(path) - 1; i < j; i, j = i + 1, j - 1 {
		path[i], path[j] = path[j], path[i]
	}
	return
}
