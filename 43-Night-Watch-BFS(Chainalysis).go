/*
You are on a night watch in a mall, represented as an n x m grid:

'x' denotes areas blocked by walls, fountains, etc., and are not traversable.
'o' denotes open spaces that the security guard can explore.
Your task is to find the shortest distances from a given starting position to every other reachable point in the grid.

eg.

[['o', 'o', 'x', 'o'],
 ['x', 'o', 'o', 'o'],
 ['o', 'o', 'x', 'o']]

 Starting Poing  = [0, 0]

 Output

 [[1, 0, -1, 4],
 [-1, 1, 2, 3],
 [3, 2, -1, 4]]
*/


/*
Solution Uses a BFS algorithm to compute the shortest path from the start position to all reachable positions.
Outputs distances or marks positions as 'x' if they are not reachable (walls or isolated by walls).
*/
package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

// Directions arrays for moving in the grid: right, left, down, up
var dirs = []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

// BFS to find shortest path in a grid
func bfs(grid [][]rune, start Point) [][]int {
	n, m := len(grid), len(grid[0])
	distances := make([][]int, n)
	for i := range distances {
		distances[i] = make([]int, m)
		for j := range distances[i] {
			if grid[i][j] == 'x' {
				distances[i][j] = -1 // Initialize walls as -1 (unreachable)
			} else {
				distances[i][j] = -1 // Initialize others as -1 initially (unvisited)
			}
		}
	}

	queue := []Point{start}
	distances[start.x][start.y] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			nx, ny := current.x+dir.x, current.y+dir.y
			if nx >= 0 && nx < n && ny >= 0 && ny < m && grid[nx][ny] == 'o' && distances[nx][ny] == -1 {
				queue = append(queue, Point{nx, ny})
				distances[nx][ny] = distances[current.x][current.y] + 1
			}
		}
	}

	return distances
}

func main() {
	grid := [][]rune{
		{'o', 'o', 'x', 'o', 'o'},
		{'o', 'x', 'o', 'x', 'o'},
		{'o', 'o', 'o', 'o', 'x'},
	}

	start := Point{0, 1} // Starting position from your screenshot
	distances := bfs(grid, start)

	for _, row := range distances {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}
