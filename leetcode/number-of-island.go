

import "strconv"

func numIslands(grid [][]byte) int {
	numsOfIslands := 0
	visited := make(map[string]bool)
	directions := [][]int{
		{-1, 0}, {0, -1}, {0, 1}, {1, 0},
	}	
	
	for i := range len(grid) {
		for j := range len(grid[0]) {
			positionValue := string(grid[i][j])
			if (positionValue == "1" && !visited[getKey(i, j)]){
				dfs(grid, i, j, directions, visited)
				numsOfIslands++	
			}
		}
	}

	return numsOfIslands
}

func dfs(grid [][]byte, x, y int, directions [][]int, visited map[string]bool){
	key := getKey(x, y)
	visited[key] = true

	for _, direction := range directions {
		nextRow := x + direction[0]
		nextCol := y + direction[1]
		
        if (nextRow >= 0 && nextRow < len(grid) && nextCol >=0 && nextCol < len(grid[0])) {
            nextPositionValue := string(grid[nextRow][nextCol])
            nextPositionKey := getKey(nextRow, nextCol)
            if (nextPositionValue == "1" && !visited[nextPositionKey]){
                dfs(grid, nextRow, nextCol, directions, visited)
            }
        }
	} 
}

func getKey(x, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}



