package main

func graph_valid_tree(n int, edges [][]int) bool {
	if (len(edges) == 0) {
		return n == 1
	}

	adjacentList := getAdjacentList(edges)
	firstEdge := edges[0][0];
	visited := make(map[int]bool, n)

	var hasCycle = hasCycle(firstEdge, -1, adjacentList, visited)
	if (hasCycle) {
		return false
	}

	return len(visited) == n
}

func hasCycle(edge int, prev int, adjacentList map[int][]int, visited map[int]bool) bool {	
	adjacents := adjacentList[edge]
    visited[edge] = true
	for _, adjacentVertex := range adjacents {
		_, exist := visited[adjacentVertex] 
		if (exist && adjacentVertex != prev) {
				return true
		} 
        if (!exist && hasCycle(adjacentVertex, edge, adjacentList, visited)) {
				return true	
		}
	}

	return false 
}

func getAdjacentList(edges [][]int) map[int][]int {
	r := make(map[int][]int)

	for _, edge := range edges {
		firstVertex, secondVertex := edge[0], edge[1]
		r[firstVertex] = append(r[firstVertex], secondVertex)
		r[secondVertex] = append(r[secondVertex], firstVertex) 
	} 
	
	return r
}
