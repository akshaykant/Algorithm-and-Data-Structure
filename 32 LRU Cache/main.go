/*
Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity,
it should invalidate the least recently used item before inserting a new item.


The cache is initialized with a positive capacity.

Follow up:
Could you do both operations in O(1) time complexity?

Example:

LRUCache cache = new LRUCache( 2 // capacity // );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.put(4, 4);    // evicts key 1
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4
*/
package main

import (
	"container/list"
)

/*
The problem can be solved with a hashtable that keeps track of the keys and its values in the double linked list.
One interesting property about double linked list is that the node can remove itself without other reference.
In addition, it takes constant time to add and remove nodes from the head or tail.

Fast removal. Doubly linked lists let us remove and insert in constant time if we have access to a node directly. The hashtable gives us access to a node directly.

If we use a singly linked list we will need to spend O(n) time to remove a node even if we have direct reference to the node that needs to get removed.
(This is because to remove in a singly linked list we need to point nodeToDelete's previous node to nodeToDelete's next node.
Finding nodeToDelete's previous is expensive if nodeToDelete is the last node in the list.)

Important: We never need to traverse (search) the linked list because the dictionary allows us to instantly look up the location of each node!
We can add nodes as long as we keep a pointer to the tail since we only ever add to the end of the list.
*/

type kv struct{
	// key is only needed to delete entry in data map
	// when removing entry from LRU
	key   int
	value int
}

type LRUCache struct{
	capacity int
	data map[int]*list.Element //Key to Doubly Linked List Node
	hits *list.List  //doubly linked list
}

func main(){

	//constructor
	lruCache := &LRUCache{
		capacity:2,
		data: make(map[int]*list.Element),
		hits: list.New(),
		}

	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Get(1)       // returns 1
	lruCache.Put(3, 3)     // evicts key 2
	lruCache.Get(2)        // returns -1 (not found)
	lruCache.Put(4, 4)     // evicts key 1
	lruCache.Get(1)        // returns -1 (not found)
	lruCache.Get(3)        // returns 3
	lruCache.Get(4)        // returns 4
}

/*
get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
*/
func (lruCache *LRUCache) Get (key int) int{

	if element, ok := lruCache.data[key]; ok{
		//move the element to the front, as this is the recent used element
		lruCache.hits.MoveToFront(element)

		//Get the Value of the element and transform it to the struct kv, for the interface
		return element.Value.(kv).value
	}
	//If element is not there, return -1
	return -1
}

/*
put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity,
it should invalidate the least recently used item before inserting a new item.
*/
func (lruCache *LRUCache) Put(key int, value int){
	//Update the element and move to front as this is recently used
	if element, ok := lruCache.data[key]; ok{

		//Update the value of the element
		element.Value = kv{key, value}

		//move to front
		lruCache.hits.MoveToFront(element)

		//exit, as not need to check the capacity because this is update
		return
	}

	//check if capacity is reached, and add the element to the Front of the List and the element to the Map
	if lruCache.hits.Len() == lruCache.capacity{
		back := lruCache.hits.Back()

		//delete the element from Map for the key of the last element in the list
		delete(lruCache.data, back.Value.(kv).key)

		//remove the back element from the list
		lruCache.hits.Remove(back)
	}

	//Put : add the element to the front of the list and to the map
	lruCache.data[key] = lruCache.hits.PushFront(kv{key: key, value: value})
}