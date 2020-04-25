/*
Given a 2D grid, each cell is either a zombie 1 or a human 0. Zombies can turn adjacent (up/down/left/right) human beings into zombies every hour. Find out how many hours does it take to infect all humans?

Example:

Input:
[[0, 1, 1, 0, 1],
 [0, 1, 0, 1, 0],
 [0, 0, 0, 0, 1],
 [0, 1, 0, 0, 0]]

Output: 2

Explanation:
At the end of the 1st hour, the status of the grid:
[[1, 1, 1, 1, 1],
 [1, 1, 1, 1, 1],
 [0, 1, 0, 1, 1],
 [1, 1, 1, 0, 1]]

At the end of the 2nd hour, the status of the grid:
[[1, 1, 1, 1, 1],
 [1, 1, 1, 1, 1],
 [1, 1, 1, 1, 1],
 [1, 1, 1, 1, 1]]
*/
package main

import "fmt"

type Index struct {
	x int
	y int
}

type Queue struct{
	mem []Index
	front int
	back int
}

func main(){

	grid := [][]int{
		{0,1,1,0,1},
		{0,1,0,1,0},
		{0,0,0,0,1},
		{0,1,0,0,0},
	}

	/*grid := [][]int{
		{1,1,1,1,1},
		{1,1,1,1,1},
		{1,1,1,1,1},
		{1,1,1,1,1},
	}
*/
	q :=  Queue{make([]Index, 0), 0, 0}

	hours := q.zombie(grid)

	fmt.Println(hours)
}

/*
Iterate though each index and check if it is a Zombie(1) or Human(0).
Add all the zombies into the queue.
Iterate through each zombie and make it move in all possible directions.
Convert human to zombie, and these zombies will become active for next hour
to convert adjacent humans into zombies.
So in a BFS manner, add these zombies to the queue and increase number of days
when all the possible humans are converted into zombies.
*/
func (q *Queue) zombie(grid [][]int) int{
	if len(grid) == 0{
		return 0
	}

	humanCount := 0

	//Iterate though each index and add to the queue, if Zombie or increase the count of the humans
	for row := 0; row < len(grid); row +=1{
		for col := 0; col < len(grid[0]); col += 1{
			if grid[row][col] == 1{
				//add to the queue
				q.Enqueue(Index{row, col})

			}else{
				humanCount += 1
			}
		}
	}

	//check if there are no humans
	if humanCount == 0{
		return 0
	}

	hours := 0

	directions := []Index{{0,1}, {0, -1}, {-1,0}, {1, 0}}

	//Iterate through each element in Queue and add newly converted zombies in a BFS way
	for !q.isEmpty() && humanCount > 0 {

		qLen := q.Len()

		for i := 0; i < qLen; i += 1{
			zombie, _ := q.Dequeue()
			for j := 0; j < len(directions); j += 1{
				newX := zombie.x + directions[j].x
				newY := zombie.y +directions[j].y

				//Check if the new index is inside the grid and is a zombie
				//Remember: to set the length of the grid - 1
				if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) && grid[newX][newY] == 0{
					//convert it into a zombie and add it to the queue so that it will be part of converting human in next hour
					grid[newX][newY] = 1
					q.Enqueue(Index{newX, newY})
					//Remember : Decrease the human count, as human is converted into a zombie
					humanCount -= 1
				}
			}
		}

		//Increment the hours
		hours += 1
	}

	return hours
}

func (q *Queue) isEmpty() bool{
	if len(q.mem) == 0{
		return true
	}
	return false
}

func (q *Queue) Len() int{
	return q.back - q.front
}

func (q *Queue) Enqueue(index Index){
	q.mem = append(q.mem, index)
	q.back = len(q.mem)
}

func (q *Queue) Dequeue()(Index, string){

	if q.front == q.back{
		return Index{}, "Error : Queue is empty"
	}

	index :=  q.mem[0]
	q.mem = q.mem[1:]
	q.back = len(q.mem)

	return index, ""

}