/*
There are n servers numbered from 0 to n-1 connected by undirected server-to-server connections forming a network
where connections[i] = [a, b] represents a connection between servers a and b. Any server can reach any other server directly or indirectly through the network.

A critical connection is a connection that, if removed, will make some server unable to reach some other server.

Return all critical connections in the network in any order.

Example 1:
	 |------2
1----|      |
|    |---0--|
3

Input: n = 4, connections = [[0,1],[1,2],[2,0],[1,3]]
Output: [[1,3]]
Explanation: [[3,1]] is also accepted.


Constraints:

1 <= n <= 10^5
n-1 <= connections.length <= 10^5
connections[i][0] != connections[i][1]
There are no repeated connections.
*/
package main

import (
	"fmt"
	"math"
)

type Node struct{
	item int
	neighbor []int
}

type Graph struct{
	mem []*Node
	time int
}

func main(){
	connections := [][]int{{0,1}, {1,2},{2,0},{1,3}}

	n := 4
	graph := &Graph{}

	//Construct the graph for the number of nodes, adding edges for both nodes
	graph.constructGraph(n, connections)

	/*for i, v := range graph.mem{
		fmt.Println(i, v.item, v.neighbor)
	}*/
	result := graph.connectedComponentTarjan(n)

	fmt.Println("Tarjan", result)

}
/*
Tarjan's Algorithm
https://iq.opengenus.org/tarjans-algorithm/
https://www.youtube.com/watch?v=erlX-1MJlv8
https://www.youtube.com/watch?v=2kREIkF9UAs

A directed graph is called strongly connected if each vertex of the graph is reachable from every other vertex in the graph.
This means that a path would exist between each pair of nodes.

For a directed graph, a strongly connected component is a partition or a subgraph, which is not a subgraph of another strongly connected component.
This means that a strongly connected component has to be the maximal subgraph satisfying the condition.
Henceforth, we will refer to strongly connected components using the abbreviated term SCC.

The steps involved are:

- A dfs is run over the nodes and the subtrees of SCCs are removed and recorded as they are encountered.

- Two values dfs_num(u) and dfs_low(u) are maintained for each of the users. dfs_num(u) is the value of the counter when the node u is explored for the first time.
dfs_low(u) stores the lowest dfs_num reachable from u which is not the part of another SCC.

- As the nodes are explored, they are pushed onto a stack.

- The unexplored children of a node are explored and dfs_low(u) is accordingly updated.

- A node is encountered with dfs_low(u) == dfs_num(u) is the first explored node in its strongly connected component and
all the nodes above it in the stack are popped out and assigned the appropriate SCC number.

Time complexity: The algorithm is built upon DFS and therefore, each node is visited once and only once.
For each node, we perform some constant amount of work and iterate over its adjacency list. Thus, the complexity is O(|V|+ |E|)

At maximum, the depth of recursion and the size of stack can be n nodes. Thus the complexity is O(|V|)
*/
func (graph *Graph) connectedComponentTarjan(n int)[][]int{

	result  := make([][]int, 0)
	dfs_low := make([]int, n)

	dfs_num := make([]int,n)

	//Mark each node to be not visited(-1)
	for i := 0; i < n; i += 1{
		dfs_num[i] = -1
	}

	//time to mark the visited nodes
	graph.time = 0

	for u := 0; u < n; u += 1{
		if dfs_num[u] == -1{
			result = graph.dfsTarjan(u, dfs_num, dfs_low, result, u)
		}
	}
	return result
}


func (graph *Graph) dfsTarjan(item int, dfs_num []int, dfs_low []int, result [][]int, parent int) [][]int{
	//When the Node is discovered
	//increment the time of dfs_num, which is when this node is visited and dfs_low for maintaining it to be part of the  connected components
	graph.time += 1
	dfs_num[item] = graph.time
	dfs_low[item] = graph.time

	//iterate through each neighbor and place them on stack by incrementing the visited time.
	for _, v := range graph.mem[item].neighbor {

		neighbor := v

		//When parent node, ignore
		if neighbor == parent {
			continue
		}
		//if not discovered
		if dfs_num[neighbor] == -1 {
			//Add it to the Stack
			result = graph.dfsTarjan(neighbor, dfs_num, dfs_low, result, item)

			dfs_low[item] = int(math.Min(float64(dfs_low[item]), float64(dfs_low[neighbor])))

			// u(item) - v(neighbor) is critical, there is no path for v to reach back to u or previous vertices of u
			//Neighbor's visited(dfs_num), if part of cycle, need to be less that the current item, as during backtracking
			//when parent node is found, we are updating the low value to that of parents visited, which will be the least.
			//So if there is cycle this condition will be satisfied else if wont be
			if dfs_low[neighbor] > dfs_num[item] {
				result = append(result, []int{item, neighbor})
			}
		} else {
			// if v discovered and is not parent of u, update low[u], cannot use low[v] because u is not subtree of v
			//This is where we are updating the low of the item, when cycle is met.
			dfs_low[item] = int(math.Min(float64(dfs_low[item]), float64(dfs_num[neighbor])))
		}
	}
	return result
}

/*
https://leetcode.com/problems/critical-connections-in-a-network/discuss/382638/No-TarjanDFS-detailed-explanation-O(orEor)-solution-(I-like-this-question)
Thinking for a little while, you will easily find out this theorem on a connected graph:

**An edge is a critical connection, if and only if it is not in a cycle.**
So, if we know how to find cycles, and discard all edges in the cycles, then the remaining connections are a complete collection of critical connections.

**How to find edges in cycles, and remove them?**
We will use DFS algorithm to find cycles and decide whether or not an edge is in a cycle.

Define rank of a node: The depth of a node during a DFS. The starting node has a rank 0.

Only the nodes on the current DFS path have non-special ranks. In other words, only the nodes that we've started visiting,
but haven't finished visiting, have ranks. So 0 <= rank < n.

(For coding purpose, if a node is not visited yet, it has a special rank -2; if we've fully completed the visit of a node, it has a special rank n.)

How can "rank" help us with removing cycles? Imagine you have a current path of length k during a DFS.
The nodes on the path has increasing ranks from 0 to k and incrementing by 1. Surprisingly, your next visit finds a node that has a rank of p
where 0 <= p < k. Why does it happen? Aha! You found a node that is on the current search path! That means, congratulations, you found a cycle!

But only the current level of search knows it finds a cycle. How does the upper level of search knows, if you backtrack?
Let's make use of the return value of DFS: dfs function returns the minimum rank it finds. During a step of search from node u to its neighbor v,
if dfs(v) returns something smaller than or equal to rank(u), then u knows its neighbor v helped it to find a cycle back to u or u's ancestor.
So u knows it should discard the edge (u, v) which is in a cycle.

After doing dfs on all nodes, all edges in cycles are discarded. So the remaining edges are critical connections.

**Complexity analysis**
DFS time complexity is O(|E| + |V|), attempting to visit each edge at most twice. (the second attempt will immediately return.)
As the graph is always a connected graph, |E| >= |V|.

So, time complexity = O(|E|).

Space complexity = O(graph) + O(rank) + O(connections) = 3 * O(|E| + |V|) = O(|E|).

FAQ: Are you reinventing Tarjan?
Honestly, I didn't know Tarjan beforehand. The idea of using rank is inspired by pre-ordering which is a basic concept of DFS.
Now I realize they are similar, but there are still major differences between them.

This solution uses only one array rank. While Tarjan uses two arrays: dfn and low.
This solution's min_back_depth is similar to Tarjan's low, but rank is very different than dfn. max(dfn) is always n-1, while max(rank) could be smaller than n-1.
This solution constructs the result by removing non-critical edges during the dfs, while Tarjan constructs the result by collecting non-critical edges after the dfs.
In this solution, only nodes actively in the current search path have 0<=rank[node]<n; while in Tarjan,
nodes not actively in the current search path may still have 0<=dfn[node]<=low[node]<n.
*/

func (graph *Graph) criticalConnection(){

}

func (graph *Graph) constructGraph(n int, connections [][]int){

	//Add all the nodes
	for i := 0; i < n; i += 1{
		graph.addNode(i)
	}

	//Add all the connections
	for _, v := range connections{
		//for undirected graph, add for both the nodes
		graph.addNeighbor(v[0], v[1])
		graph.addNeighbor(v[1], v[0])

	}
}


func (graph *Graph) addNode(item int){
	graph.mem = append(graph.mem, &Node{item, nil})
}

func (graph *Graph) addNeighbor(item int, neighbor int){

		if graph.mem[item].item == item {
			graph.mem[item].neighbor = append(graph.mem[item].neighbor, neighbor)
		}
}