/*
Given a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

Example:

Input: [1,2,3,null,5,null,4]
Output: [1, 3, 4]
Explanation:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
*/

package main

import "fmt"

type Node struct{
	item int
	left *Node
	right *Node
}

type Queue struct{
	list []*Node
}

func main(){

	leaf1 := Node{5, nil, nil}

	leaf2 := Node{4, nil, nil}

	lev1 := Node{2, nil, &leaf1}

	lev2 := Node{3, nil, &leaf2}

	root := Node{1, &lev1, &lev2}


	q := Queue{}

	out := q.rightSideView(&root, []int{})

	fmt.Println(out)
}

func (q *Queue)rightSideView(root *Node,visibleList []int) []int{

	if root == nil{
		return visibleList
	}

	q.Enqueue(root)

	//Until Queue is empty
	for !q.isEmpty() {

		size := q.size()

		for i := 0; i < size; i += 1{
			current := q.Dequeue()

			//last node for a level
			if i == size - 1{
				visibleList = append(visibleList, current.item)
			}
			if current.left != nil{
				q.Enqueue(current.left)
			}
			if current.right != nil{
				q.Enqueue(current.right)
			}
		}
	}

	return visibleList
}

func (q *Queue)Enqueue(node *Node){
	//enqueue at end
	q.list = append(q.list, node)
}

func (q *Queue)Dequeue()*Node{

	if len(q.list) == 0{
		return nil
	}
	//dequeue from start
	node := q.list[0]

	q.list = q.list[1:]

	return node
}

func (q *Queue)isEmpty() bool{

	if len(q.list) == 0{
		return true
	}
	return false
}

func (q *Queue)size()int{
	return len(q.list)
}