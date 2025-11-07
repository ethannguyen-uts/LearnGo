package main

import "strconv"

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	originColor := image[sr][sc]
	directions := [][]int{
		{0, -1},
		{-1, 0},
		{0, 1},
		{1, 0},	
	} 

	dfsHelper(image, sr, sc, color, originColor, directions, make(map[string]bool))
	
	return image
}

func dfsHelper(image [][]int, r int, c int, color int, originColor int, directions [][]int, visited map[string]bool){
	key := getKey(r, c)
	visited[key] = true
	image[r][c] = color

	maxRow := len(image)
	maxCol := len(image[0])
	
	for _, direction := range directions {
		nextRow := r + direction[0]
		nextCol := c + direction[1]
		
		if (nextRow < 0 || nextRow >= maxRow || nextCol < 0 || nextCol >= maxCol){
			continue
		}

        nextKey := getKey(nextRow, nextCol)
		if (image[nextRow][nextCol] == originColor && !visited[nextKey]){
			dfsHelper(image, nextRow, nextCol, color, originColor, directions, visited)
		}
	} 
		
}

func getKey(r, c int) string {
    return strconv.Itoa(r) + "," + strconv.Itoa(c)
}
