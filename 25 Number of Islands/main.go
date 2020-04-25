/*
Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water
and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:

Input:
11110
11010
11000
00000

Output: 1

Example 2:

Input:
11000
11000
00100
00011

Output: 3
*/
package main

import "fmt"

func main(){


/*	grid := [][]int{
		{1,1,1,1,0},
		{1,1,0,1,0},
		{1,1,0,0,0},
		{0,0,0,0,0}}*/

	grid := [][]int{
				{1,1,0,0,0},
				{1,1,0,0,0},
				{0,0,0,0,0},
				{0,0,1,0,0},
				{0,0,0,1,1}}

	islands := numberOfIslands(grid)

	fmt.Println(islands)
}
//Iterate through each point and if land(1), recursively check all the neighbors and mark them as Water(0) if land.
//Neighbors are horizontal or vertical points. Diagonal points are not referred as neighbors.
func numberOfIslands(grid [][]int) int{

	islands := 0
	if len(grid) == 0{
		return islands
	}

	//Iterate through each element in the grid to check if it is land(1) or water(0)
	for row := 0; row < len(grid); row += 1{
		for column := 0; column < len(grid[0]); column += 1{
			if grid[row][column] == 1{

				dfs(grid, row, column)
				islands += 1
			}
		}
	}
	return islands
}
/*
Marks the given site as visited, then checks adjacent sites.
Or, Marks the given site as water, if land, then checks adjacent sites.

Or, Given one coordinate (row,col) of an island, obliterates the island
from the given grid, so that it is not counted again.
*/
func dfs(grid [][]int, row int, col int){
	//Sink the site - mark land to water.
	grid[row][col] = 0

	//Check all the adjacent sites, to check if they are land and are valid indices
	if row - 1 >= 0 && grid[row - 1][col] == 1{
		dfs(grid, row - 1, col)
	}
	if row + 1 < len(grid) && grid[row + 1][col] == 1{
		dfs(grid, row + 1, col)
	}
	if col - 1 >= 0 && grid[row][col - 1] == 1{
		dfs(grid, row, col - 1)
	}
	if col + 1 < len(grid[0]) && grid[row][col + 1] == 1{
		dfs(grid, row, col + 1)
	}
}