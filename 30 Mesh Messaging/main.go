/*
You wrote a trendy new messaging app, MeshMessage, to get around flaky cell phone coverage.

Instead of routing texts through cell towers, your app sends messages via the phones of nearby users,
passing each message along from one phone to the next until it reaches the intended recipient. (Don't worry—the messages are encrypted while they're in transit.)

Some friends have been using your service, and they're complaining that it takes a long time for messages to get delivered.
After some preliminary debugging, you suspect messages might not be taking the most direct route from the sender to the recipient.

Given information about active users on the network, find the shortest route for a message from one user (the sender) to another (the recipient).
Return an array of users that make up this route.

There might be a few shortest delivery routes, all with the same length. For now, let's just return any shortest route.

Your network information takes the form of a hash map mapping username strings to an array of other users nearby:

Example 1:

 network = {
  'Min'   : ['William', 'Jayden', 'Omar'],
  'William' : ['Min', 'Noam'],
  'Jayden' : ['Min', 'Amelia', 'Ren', 'Noam'],
  'Ren'   : ['Jayden', 'Omar'],
  'Amelia' : ['Jayden', 'Adam', 'Miguel'],
  'Adam'  : ['Amelia', 'Miguel', 'Sofia', 'Lucas'],
  'Miguel' : ['Amelia', 'Adam', 'Liam', 'Nathan'],
  'Noam'  : ['Nathan', 'Jayden', 'William'],
  'Omar'  : ['Ren', 'Min', 'Scott'],
  ...
}

For the network above, a message from Jayden to Adam should have this route:
 ['Jayden', 'Amelia', 'Adam']
*/

package main

import "fmt"

type Node struct {
	item     string
	neighbor []string
}

type Graph struct {
	mem map[string]*Node
}

type Queue struct {
	mem   []string
	front int
	back  int
}

func main() {

	memory := make(map[string]*Node)

	graph := &Graph{memory}

	graph.addNode("Min")
	graph.addNode("William")
	graph.addNode("Jayden")
	graph.addNode("Ren")
	graph.addNode("Amelia")
	graph.addNode("Adam")
	graph.addNode("Miguel")
	graph.addNode("Noam")
	graph.addNode("Omar")
	graph.addNode("Sofia")
	graph.addNode("Lucas")
	graph.addNode("Liam")
	graph.addNode("Nathan")
	graph.addNode("Scott")
	graph.addNeighbor("Min", []string{"William", "Jayden", "Omar"})
	graph.addNeighbor("William", []string{"Min", "Noam"})
	graph.addNeighbor("Jayden", []string{"Min", "Amelia", "Ren", "Noam"})
	graph.addNeighbor("Ren", []string{"Jayden", "Omar"})
	graph.addNeighbor("Amelia", []string{"Jayden", "Adam", "Miguel"})
	graph.addNeighbor("Adam", []string{"Amelia", "Miguel", "Sofia", "Lucas"})
	graph.addNeighbor("Miguel", []string{"Amelia", "Adam", "Liam", "Nathan"})
	graph.addNeighbor("Noam", []string{"Nathan", "Jayden", "William"})
	graph.addNeighbor("Omar", []string{"Ren", "Min", "Scott"})

	res := graph.communicationPath("Jayden", "Adam")

	fmt.Println(res)

}

/*
Since we're interested in finding the shortest path, BFS is the way to go.

Remember: both BFS and DFS will eventually find a path if one exists. The difference between the two is:
- BFS always finds the shortest path.
- DFS usually uses less space.

For every node in the graph, I will start from the start node and will keep on traversing through every node
via BFS, adding each neighbor to the queue, until the destination node is reached.

For keeping track of the path,
*/
func (graph *Graph) communicationPath(from string, to string) []string {

	/*
		Look at the nodes_already_seen set—that's really important and easy to forget. If we didn't have it,
		our algorithm would be slower (since we'd be revisiting tons of nodes) and it might never finish (if there's no path to the end node).
	*/
	nodeAlreadyVisited := make(map[string]bool)

	/*
		We're using a queue instead of a list because we want an efficient first-in-first-out
		(FIFO) structure with O(1)O(1) inserts and removes. If we used a list, appending would be O(1)O(1),
		but removing elements from the front would be O(n)O(n).
	*/
	nodeToVisit := &Queue{make([]string, 0), 0, 0}

	/*
		This is for book-keeping to store the which node is visited by which node.
		We can reconstruct our path by traversing visitPath, From to To node and reverse than to get the path.

		We'd take this dictionary we built up during our search:

		 {'Min'   : START,
		 'Jayden' : 'Min',
		 'Ren'   : 'Jayden',
		 'Amelia' : 'Jayden',
		 'Adam'  : 'Amelia',
		 'Miguel' : 'Amelia',
		 'William' : 'Min'}

		And, we'd use it to backtrack from the end node to the start node, recovering our path:

		To get to Adam, we went through Amelia.
		To get to Amelia, we went through Jayden.
		To get to Jayden, we went through Min.
		Min is the start node.
	*/
	visitPath := make(map[string]string)

	nodeToVisit.Enqueue(from)
	nodeAlreadyVisited[from] = true
	visitPath[from] = "START"

	//BFS: which will traverse through all the neighbors first and then their child
	for !nodeToVisit.isEmpty() {

		current, _ := nodeToVisit.Dequeue()

		neighbor := graph.mem[current].neighbor

		//found the path
		if current == to {

			return reversePath(visitPath, to, from)
		}

		for _, v := range neighbor {
			if !nodeAlreadyVisited[v] {
				nodeAlreadyVisited[v] = true
				nodeToVisit.Enqueue(v)

				visitPath[v] = current

			}
		}
	}
	//When traversed through all the nodes and no path is found
	return nil
}

func reversePath(visitedPath map[string]string, to string, from string) []string {

	var res []string

	res = append(res, to)
	//traverse from last to first
	current := visitedPath[to]

	for !(current == "START") {

		res = append(res, current)

		current = visitedPath[current]
	}

	// reverse the results path
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[j], res[i] = res[i], res[j]
	}

	return res
}

/*
For directed vs. undirected, we'll assume that if Min can transmit a message to Jayden,
then Jayden can also transmit a message to Min. Our sample input definitely suggests this is the case.
And it makes sense—they're the same distance from each other, after all. That means our graph is undirected.

What about weighted? We're not given any information suggesting that some transmissions are more expensive than others, so let's say our graph is unweighted.
*/
func (graph *Graph) addNode(item string) {
	graph.mem[item] = &Node{item, nil}
}

func (graph *Graph) addNeighbor(item string, neighbor []string) {

	//Range though all the neighbors and add item to each neighbor and neighbor to each item, for bi-directional
	for _, v := range neighbor {
		graph.mem[item].neighbor = append(graph.mem[item].neighbor, v)
		graph.mem[v].neighbor = append(graph.mem[v].neighbor, item)
	}
}

func (q *Queue) isEmpty() bool {
	if q.front == q.back {
		return true
	}
	return false
}

func (q *Queue) Enqueue(item string) {
	q.mem = append(q.mem, item)
	q.back += 1
}

func (q *Queue) Dequeue() (string, string) {
	if q.isEmpty() {
		return "", "Error : Empty Queue, cannot delete"
	}

	item := q.mem[q.front]

	q.mem = q.mem[1:]

	q.back = len(q.mem)

	return item, ""
}
