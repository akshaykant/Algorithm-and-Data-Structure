/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/
package main

import "fmt"

//Linked List Struct
type Node struct{
	value int
	next *Node
}

func main (){

	node1 := &Node{
		4,
		nil,
	}
	node2 := &Node{2,
		node1,
	}

	node3 := &Node{3,
		node1,
	}

	input1 := &Node{
		1,
		node2,
	}

	input2 := &Node{
		1,
		node3,
	}

	output := merge2SortedList(input1, input2)

	for output != nil {

		fmt.Println(output.value)
		output = output.next

	}


}

func merge2SortedList(list1 *Node, list2 *Node) *Node{

	//Iterate both the list and check which element is minimum. Increment the pointer of hte minimum list and store the value in another list

	var(
		mergedList = Node{}
		curr = &mergedList
	)

	for list1 != nil || list2 != nil{

		n := Node{}

		switch {

		case list1 != nil && list2 != nil:
			//check for min node
			if list1.value < list2.value {
				n.value = list1.value
				list1 = list1.next
				break
			}
			n.value = list2.value
			list2 = list2.next

		case list1 != nil && list2 == nil:
			n.value = list1.value
			list1 = list1.next

		case list1 == nil && list2 != nil:
			n.value = list2.value
			list2 = list2.next
		}

		curr.next = &n

		curr = curr.next



	}
	return mergedList.next
}