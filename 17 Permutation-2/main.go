/*
Given a collection of duplicate integers, return all possible permutations.

Example:

Input: aab
Output:
[
 aab,
 aba,
 baa,
]
*/

package main

import "fmt"

type Memory struct {
	list map[string]string
}

func main() {

	input := "aab"

	m := make(map[string]string)
	mem := Memory{
		m,
	}

	mem.permute(input, 0, len(input))

	for _, v := range mem.list {
		fmt.Println(v)
	}

}

//RuneTime : O(N!)
//SpaceTime : O(N!)
//This is the recursive function which take three parameters - string, starting index and end index.
func (m *Memory) permute(str string, start int, end int) {
	//if start and end index are same : we have reached the end computation. This is the base case
	if start == end {
		m.list[str] = str
	}
	//Permute all possible solution by recursion and backtracking the entire knowledge graph of string.
	for i := start; i < end; i += 1 {
		swapStr := swap(str, start, i)
		m.permute(swapStr, start+1, end)
	}

}

//Swap the characters in the string by taking inputs as string and two index two swap
func swap(str string, i int, j int) string {

	runes := []rune(str)

	runes[i], runes[j] = runes[j], runes[i]

	return string(runes)
}
