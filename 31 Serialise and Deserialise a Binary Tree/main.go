/*
https://www.youtube.com/watch?v=suj1ro8TIVY

Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer,
or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work.
You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

Example:

You may serialize the following tree:

  1
 / \
 2  3
   / \
  4  5

as pre-order (rootLeftRight)
"[1,2, null, null,3, 4, null,null,5, null, null]"
*/

package main

import "fmt"

type Node struct {
	item  string
	left  *Node
	right *Node
}

type Queue struct {
	mem   []string
	front int
	back  int
}

type Memory struct {
	mem []string
}

func main() {

	queue := Queue{make([]string, 0), 0, 0}

	in := []string{"1", "2", "X", "X", "3", "4", "X", "X", "5", "X", "X"}

	for _, v := range in {
		queue.Enqueue(v)
	}
	//Place all the nodes in the queue and use them to create the Node with left and right child.
	//If the node is null(X), that is the end of the Node
	deserialised := queue.deserialize()

	//Memory to store deserialize data
	res := &Memory{make([]string, 0)}
	res.serialize(deserialised)

	fmt.Println(res)
}

func (memory *Memory) serialize(root *Node) {

	//When we encounter no leaf, we attach X
	if root == nil {
		memory.mem = append(memory.mem, "X")
		//remember to return, else this base condition won't be return
		return
	}

	memory.mem = append(memory.mem, root.item)
	memory.serialize(root.left)
	memory.serialize(root.right)

}

func (queue *Queue) deserialize() *Node {

	current, _ := queue.Dequeue()
	//
	//If current node is Null, that is end of the Node. Next item will be added to the parent of it.
	if current == "X" {
		return nil
	}

	//The last string element in the queue must be "X", because in the serialization process,
	// we represent each null node (including the child nodes of leaf node) as "X".
	//When deserializing, the function returns null whenever meeting a "X" from queue.
	// So when the queue is null, it will return null directly.
	//if(current == nil) return nil  //this is not needed

	node := &Node{current, nil, nil}

	node.left = queue.deserialize()
	node.right = queue.deserialize()

	return node
}

func (queue *Queue) isEmpty() bool {
	if queue.front == queue.back {
		return true
	}
	return false
}

func (queue *Queue) Enqueue(item string) {

	queue.mem = append(queue.mem, item)
	queue.back += 1
}

func (queue *Queue) Dequeue() (string, string) {

	if queue.isEmpty() {
		return "", "Error: empty queue"
	}

	item := queue.mem[queue.front]

	queue.mem = queue.mem[1:]

	queue.back = len(queue.mem)

	return item, ""
}
