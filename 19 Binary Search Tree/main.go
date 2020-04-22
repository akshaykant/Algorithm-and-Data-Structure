package main

import "fmt"

type Node struct {
	item int
	left *Node
	right *Node
}

type Memory struct {
	list []int
}
func main(){

	values := []int{23, 54, 12, 876, 43, 32, 686, 434, 56, 4, 1, 14}

	var root *Node

	for _, v := range  values{
		root = insert(v, root)
		fmt.Println(v, root)
	}

	m := Memory{}
	//left Root right
	fmt.Println("left Root right")
	m.traverseInOrder(root)

	fmt.Println(m)

	//Root Left Right
	fmt.Println("Root Left Right")
	traversePreOrder(root)

	//left right Root
	fmt.Println("left right Root")
	traversePostOrder(root)

	//Search
	fmt.Println("Search")
	b := search(876, root)
	fmt.Println(b)

	fmt.Println("Min")
	fmt.Println(min(root))

	fmt.Println("Max")
	fmt.Println(max(root))

	root = remove(876, root)
	fmt.Println("Remove")
	//traverseInOrder(root)

	fmt.Println("Height")
	fmt.Println(height(root))
}



func insert(item int, root *Node)*Node{

	if root == nil{
		root = &Node{item, nil, nil}
	}
	//insert left
	if root.item > item{
		root.left = insert(item, root.left)
	}
	//insert right
	if root.item < item{
		root.right = insert(item, root.right)
	}

	return root
}

func (m *Memory) traverseInOrder(root *Node){

	if root == nil{
		return
	}
	if root.left != nil{
		m.traverseInOrder(root.left)
	}
	//fmt.Println(root.item)
	m.list = append(m.list, root.item)
	if root.right != nil {
		m.traverseInOrder(root.right)
	}
}

func traversePreOrder(root *Node){

	if root == nil{
		return
	}
	fmt.Println(root.item)
	if root.left != nil{
		traversePreOrder(root.left)
	}
	if root.right != nil{
		traversePreOrder(root.right)
	}
}

func traversePostOrder(root *Node){
	if root == nil{
		return
	}
	if root.left != nil{
		traversePostOrder(root.left)
	}
	if root.right != nil{
		traversePostOrder(root.right)
	}
	fmt.Println(root.item)
}

func search(value int, root *Node) bool{
	if root == nil{
		return false
	}
	if root.item == value{
		return true
	}
	if root.item > value{
		return search(value, root.left)
	}
	if root.item < value{
		return search(value, root.right)
	}
	return false
}

func min(root *Node) int{

	if root == nil{
		return -1
	}

	if root.left == nil{
		return root.item
	}

	return min(root.left)
}

func max(root *Node)int{

	if root == nil{
		return -1
	}

	if root.right == nil{
		return root.item
	}


	return max(root.right)

}

func remove(item int, root *Node) *Node{

	if root == nil{
		return root
	}
	if root.item > item{
		root.left = remove(item, root.left)
	}
	if root.item < item{

		root.right = remove(item, root.right)
	}

	if root.item == item{
		//no child
		if root.left == nil && root.right == nil{
			root = nil
			return nil
		}
		//one child
		if root.right == nil{
			root =  root.left
		}
		if root.left == nil{
			root = root.right
		}

		//both child
		if root.left != nil && root.right != nil{
			root.item = min(root.right)
			root.right = remove(root.item, root.right)
		}

	}
	return root
}
/*

The height of a tree is the length of the path from the root to the deepest node in the tree. A (rooted) tree with only a node (the root) has a height of zero.
If there is no node, you want to return -1 not 0. This is because you are adding 1 at the end.

So if there isn't a node, you return -1 which cancels out the +1.
*/
func height(root *Node) int{
	if root == nil{
		return -1
	}

	lefth := height(root.left)
	righth := height(root.right)

	if lefth > righth {
		return lefth + 1
	} else {
		return righth + 1
	}
}