package main

import "fmt"

func createAdjecentList(edges [][]int) map[int][]int {
	r := make(map[int][]int)

	for i := range edges {
		first, second := edges[i][0], edges[i][1]
		firstVertexNeigbours, ok := r[first]
		if (ok) {
			r[first] = append(firstVertexNeigbours, second)
		} else {
			r[first] = []int{second}
		}

		secondVertexNeigbours, ok := r[second]

		if (ok) {
			r[second] = append(secondVertexNeigbours, first)
		} else {
			r[second] = []int{first}
		}
	}

	return r
}

func cCreateAdjacentListMain() {
	edges := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 0},
		{0, 2},
	}

	fmt.Println(createAdjecentList(edges))
}
