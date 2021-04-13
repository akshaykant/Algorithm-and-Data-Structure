//Apple
/**
Picture the scene, you're surfing GitHub,
and would like to find a file in your repo.
You click the button 'Find File' which allows
to search for all the fully qualified file paths that match your query.
E.g if you have the following 3 files:
'/a/b/c/d.java'
'/a/b/e.java'
'/a/p/q/r.java'

then the search query '/a/b' should return these 2 files:
'/a/b/c/d.java'
'/a/b/e.java'

Think of a data structure, and an algorithm to perform on that data structure, well suited to such a search.
And now implement both :)

//
);// List<String> path1 = Arrays.asList("MyModule", "src", "com", "company", "hello", "Hello.java");
///MyModule/sre/com/company/hello/hello.java
//// List<String> path2 = Arrays.asList("MyModule", "src", "com", "company", "hello", "utils", "HelloUtils.java");
//// List<String> path3 = Arrays.asList("MyModule", "src", "com", "company", "hi", "Hi.java");
//// List<String> path4 = Arrays.asList("YourModule", "src", "com", "company", "hi", "Hi.java"
// GoodDataStructure goodDataStructure = new GoodDataStructure(Arrays.asList(path1, path2, path3, path4));
// /**********/

// /* TEST CASE 1 */
// List<String> testPath1 = Arrays.asList("MyModule", "src", "com", "company", "hello", "Hello.java");
// Set<List<String>> expectedMatches1 = Set.of(path1);
// testFindMatches(goodDataStructure, testPath1, expectedMatches1, 1);
// /***************/

// /* TEST CASE 2 */
// List<String> testPath2 = Arrays.asList("MyModule", "src", "com", "company");
// Set<List<String>> expectedMatches2 = Set.of(path1, path2, path3);
// testFindMatches(goodDataStructure, testPath2, expectedMatches2, 2);
// /***************/

// /* TEST CASE 3 */
// List<String> testPath3 = Arrays.asList("YourModule", "src", "com", "company");
// Set<List<String>> expectedMatches3 = Set.of(path4);
// testFindMatches(goodDataStructure, testPath3, expectedMatches3, 3);
// /***************/

// /* TEST CASE 4 */
// List<String> testPath4 = Arrays.asList("MyModule", "src", "com", "hi", "Hi.java");
// Set<List<String>> expectedMatches4 = Collections.emptySet();
// testFindMatches(goodDataStructure, testPath4, expectedMatches4, 4);
// /***************/

// /* TEST CASE 5 */
// List<String> testPath5 = Collections.emptyList();
// Set<List<String>> expectedMatches5 = Set.of(path1, path2, path3, path4);
// testFindMatches(goodDataStructure, testPath5, expectedMatches5, 5);
// /***************/
//}

package main

import "fmt"

type Node struct {
	value string
	child []*Node
}

func main() {

	path1 := []string{"MyModule", "src", "com", "company", "hello", "Hello.java"}
	path2 := []string{"MyModule", "src", "com", "company", "hello", "utils", "HelloUtils.java"}
	path3 := []string{"MyModule", "src", "com", "company", "hi", "Hi.java"}
	path4 := []string{"YourModule", "src", "com", "company", "hi", "Hi.java"}

	root := insert(path1, nil)
	root = insert(path2, root)
	root = insert(path3, root)
	root = insert(path4, root)

	fmt.Print(root)

}

func insert(path []string, root *Node) *Node {
	if len(path) == 0 {
		return root
	}

	//condition to add root node
	if root == nil {
		node := &Node{}
		child := &Node{}
		node.value = "/"
		child.value = path[0]
		node.child = append(node.child, child)

		child = insert(path[1:], child)
		return node
	}

	//check if there are no children
	if len(root.child) == 0 {
		child := &Node{}
		child.value = path[0]
		root.child = append(root.child, child)
		child = insert(path[1:], child)
		return root
	}

	//if there are children
	present, item := isPresent(root.child, path[0])

	//if child is not present, add it to the list
	if !present {
		child := &Node{}
		child.value = path[0]
		root.child = append(root.child, child)
		child = insert(path[1:], child)
		return root
	} else {
		item = insert(path[1:], item)
		return root
	}

}

func search(path []string, root *Node) {

}

func isPresent(list []*Node, item string) (bool, *Node) {
	for _, v := range list {
		if v.value == item {
			return true, v
		}
	}
	return false, nil
}
