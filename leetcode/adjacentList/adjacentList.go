package main

import "fmt"

func createAdjecentList(edges [][]int) map[int][]int {
	r := make(map[int][]int)

	for _, edge := range edges {
		firstVertex, secondVertex := edge[0], edge[1]
		r[firstVertex] = append(r[firstVertex], secondVertex)
		r[secondVertex] = append(r[secondVertex], firstVertex)
	}

	return r
}

func test() {
	edges := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 0},
		{0, 2},
	}

	fmt.Println(createAdjecentList(edges))
}
