/*https://leetcode.com/problems/advantage-shuffle/discuss/190310/Explain-the-Algorithm-and-concise-Java-Implementation  
Given two arrays A and B of equal size, the advantage of A with respect to B is the number of indices i for which A[i] > B[i].

Return any permutation of A that maximizes its advantage with respect to B.



Example 1:

Input: A = [2,7,11,15], B = [1,10,4,11]
Output: [2,11,7,15]

Example 2:

Input: A = [12,24,8,32], B = [13,25,32,11]
Output: [24,32,8,12]


Note:

1 <= A.length = B.length <= 10000
0 <= A[i] <= 10^9
0 <= B[i] <= 10^9
*/
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Item struct {
	value int
	priority int

	index int
}

type PriorityQueue []*Item

func main(){
	A := []int{2,7,11,15}
	B := []int{1,10,4,11}

	out := shuffle(A, B)

	fmt.Println(out)

}

/*
This is a generalization of the classic Chinese story of General Tian's Horse Race.
https://en.wikipedia.org/wiki/Tian_Ji

In the original story there are only 3 numbers in each array and A[0] < B[0], A[1] < B[1], A[2] < B[2], how can A make a possible win?

The key mindset, is, instead of try to maximize the gain, try to minimize the loss.

Two Rules:
- We should first satisfy the biggest element of B, because they are the hardest to satisfy
- If the biggest value of A cannot satisfy the value of B, nothing can satisfy
*/
func shuffle(A []int, B []int)[]int{

	//Create result
	out := make([]int, len(A))

	//Sort A
	sort.Ints(A)

	//Create the Priority Queue with values of B in descending order
	pq := &PriorityQueue{}

	for i := 0; i < len(B); i = i + 1{
		heap.Push(pq, &Item{i, B[i], i})
	}

	//initialize with A - slow = 0, fast = n-1
	slow, fast := 0, len(A) - 1

	for pq.Len() > 0 {

		// If my fastest horse remained is slower than my opponents' fastest horse,
		// there is no way for me to win, use my slower horse.
		// Otherwise use my fastest horse to win this round.
		// Why using my second fastest horse won't improve my global scores?
		// If my second fastest horse is faster than my opponents' fastest one,
		// it sure is faster than the rest of his horse. Thus proved this strategy is
		// optimal.

		it := heap.Pop(pq)

		v := it.(*Item)

		//If B is greater than highest in A, use low of A else high of A.
		if v.priority >= A[fast]{
			out[v.index] = A[slow]
			slow = slow + 1
		} else {
			out[v.index] = A[fast]
			fast = fast - 1
		}
	}

	return out
}

func (pq PriorityQueue) Len() int{
	return len(pq)
}

func (pq PriorityQueue) Less( i, j int) bool{
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int){
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = pq[j].index, pq[i].index
}

func (pq *PriorityQueue) Push(x interface{}){
	item := x.(*Item)
	len := len(*pq)

	item.index = len
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{}{
	old := *pq
	len := len(old)

	item := old[len - 1]

	old[len - 1] = nil
	//item.index = -1

	*pq = old[:len-1]

	return item
}

func (pq *PriorityQueue) Update(item *Item, value int, priority int){
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}