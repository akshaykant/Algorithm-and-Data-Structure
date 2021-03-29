/*
https://hackernoon.com/today-i-learned-using-priority-queue-in-golang-6f71868902b7
https://golang.org/src/container/heap/example_pq_test.go
Given a list of reviews, a list of keywords and an integer k. Find the most popular k keywords in order of most to least frequently mentioned.

The comparison of strings is case-insensitive. If keywords are mentioned an equal number of times in reviews, sort alphabetically.

Example 1:

Input:
k = 2
keywords = ["anacell", "cetracular", "betacellular"]
reviews = [
 "Anacell provides the best services in the city",
 "betacellular has awesome services",
 "Best services provided by anacell, everyone should use anacell",
]

Output:
["anacell", "betacellular"]

Explanation:
"anacell" is occuring in 2 different reviews and "betacellular" is only occuring in 1 review.
Example 2:

Input:
k = 2
keywords = ["anacell", "betacellular", "cetracular", "deltacellular", "eurocell"]
reviews = [
 "I love anacell Best services; Best services provided by anacell",
 "betacellular has great services",
 "deltacellular provides much better services than betacellular",
 "cetracular is worse than anacell",
 "Betacellular is better than deltacellular",
]

Output:
["betacellular", "anacell"]

Explanation:
"betacellular" is occuring in 3 different reviews. "anacell" and "deltacellular" are occuring in 2 reviews, but "anacell" is lexicographically smaller.
*/
package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type Item struct {
	value    string // The value of the item; arbitrary
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of item in the heap
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func main() {

	k := 2
	/*keywords := []string{"anacell", "cetracular", "betacellular"}
	reviews := []string{"Anacell provides the best services in the city",
	"betacellular has awesome services",
	"Best services provided by anacell, everyone should use anacell",}*/

	keywords := []string{"anacell", "betacellular", "cetracular", "deltacellular", "eurocell"}
	reviews := []string{"I love anacell Best services; Best services provided by anacell",
		"betacellular has great services",
		"deltacellular provides much better services than betacellular",
		"cetracular is worse than anacell",
		"Betacellular is better than deltacellular"}

	result := frequentWord(k, keywords, reviews)

	fmt.Println(result)
}

/*
Create a map of the keyword
Iterate over each review
Initialize a Set, for each review iteration to check if the word is seen of not
Check for each work in review, if it is not seen in Set and keyword is present, add it to seen set and increment the word count in the map
Use Priority queue to hold the top k words in the priority queue and extract the top k words.

*/
func frequentWord(k int, keywords []string, reviews []string) []string {

	//Create a map of the keyword
	frequencyMap := make(map[string]int)
	for _, key := range keywords {
		frequencyMap[key] = 0
	}

	//Iterate over each review
	for _, review := range reviews {
		//Convert to Lower case
		review = strings.ToLower(review)
		/*Fields splits the string s around each instance  of one or more consecutive white space characters,
		returning an array of substrings of s or an empty list if s contains only white space.*/
		words := strings.Fields(review)

		//Initialize a Set, for each review iteration to check if the word is seen of not
		seen := make(map[string]string)

		for _, word := range words {
			_, wordSeen := seen[word]
			freq, keywordPresent := frequencyMap[word]
			//Check for each work in review, if it is not seen in Set and keyword is present, add it to seen set and increment the word count in the map
			if !wordSeen && keywordPresent {
				seen[word] = word
				frequencyMap[word] = freq + 1
			}
		}
	}

	//Use Priority queue to hold the top k words in the priority queue and extract the top k words.
	pq := &PriorityQueue{}
	i := 0
	for w, c := range frequencyMap {
		heap.Push(pq, &Item{w, c, i})
		i = i + 1
		//Only Keep k elements
		if pq.Len() > k {
			heap.Pop(pq)
			i = i - 1
		}
	}

	res := make([]string, k)
	for i := k - 1; i >= 0; i = i - 1 {
		item := heap.Pop(pq)
		wc := item.(*Item)
		res[i] = wc.value
	}

	return res
}

// pre-defined function of Sort Interface, which is a part of heap interface
func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest based on occurrence number as the priority
	// The higher the occurrence, the lower the priority. As we are storing k elements,
	// So will Pop out lower occurrence value, termed with higher priority.

	//When two elements have equal priority, need to check the length of the word. As need to store lexicographically smaller words.
	if pq[i].priority == pq[i].priority {
		return pq[i].value > pq[j].value
	}
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = pq[j].index, pq[i].index
}

// We just implement the pre-defined function in interface of heap.
func (pq *PriorityQueue) Push(x interface{}) {
	length := len(*pq)
	item := x.(*Item)
	//Add the element at the end of the array. Heapify will move it up at the right position based on index of other items
	item.index = length
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	length := len(old)
	//Swap the first element to the last element, removes the least element. Heapify will move the first element to the right position based on index of other items
	item := old[length-1]
	old[length-1] = nil // avoid memory leak
	item.index = -1     // for safety
	*pq = old[:length-1]
	return item
}

/*Usage: heap.Push(&pq, item)
  pq.update(item, item.value, 5)*/
func (pq *PriorityQueue) Update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
