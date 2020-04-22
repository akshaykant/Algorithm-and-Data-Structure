/*
Basic - https://www.youtube.com/watch?v=GCm7m5671Ps
Basics + Implementation - https://www.youtube.com/watch?v=TnZHaH9i6-0
Given a collection of distinct integers, return all possible permutations.

Example:

Input: abc
Output:
[
  abc,
  acb,
  bac,
  bca,
  cab,
  cba
]
*/
/*
Backtracking is an algorithm for finding all solutions by exploring all potential candidates.
If the solution candidate turns to be not a solution (or at least not the last one),
backtracking algorithm discards it by making some changes on the previous step, i.e. backtracks and then try again.

Here is a backtrack function which takes the index of the first integer to consider as an argument backtrack(first).

If the first integer to consider has index n that means that the current permutation is done.
Iterate over the integers from index first to index n - 1.
Place i-th integer first in the permutation, i.e. swap(nums[first], nums[i]).
Proceed to create all permutations which starts from i-th integer : backtrack(first + 1).
Now backtrack, i.e. swap(nums[first], nums[i]) back.
*/
/*
abc

i = 0 (a)

abc
acb

i = 1 (b)
bac
bca

i = 2 (c)
cab
cba
*/
package main

import "fmt"

type Memory struct{
	list []string
}

func main() {

	input := "abcde"

	mem := Memory{}

   mem.permute(input, 0, len(input))

	fmt.Println(mem.list)

}
//RuneTime : O(N!)
//SpaceTime : O(N!)
//This is the recursive function which take three parameters - string, starting index and end index.
func (m *Memory) permute(str string, start int, end int)  {
	//if start and end index are same : we have reached the end computation. This is the base case
	if start ==  end{
		m.list = append(m.list, str)
	}
	//Permute all possible solution by recursion and backtracking the entire knowledge graph of string.
	for i := start; i < end; i += 1{
		swapStr := swap(str, start, i)
		m.permute(swapStr, start + 1, end)
	}

}
//Swap the characters in the string by taking inputs as string and two index two swap
func swap(str string, i int, j int) string{

	runes := []rune(str)

	runes[i], runes[j] = runes[j], runes[i]

	return string(runes)
}