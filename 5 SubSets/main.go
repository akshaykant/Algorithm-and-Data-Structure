/*
Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

*/
package main

import "fmt"

func main(){
	input := []int{1,2,3}

	res := findSubset(input)

	fmt.Print(res)


}

func findSubset(input []int)(res [][]int){
	//add an empty list
	res = append(res, []int{})

	for i := 0; i < len(input); i = i + 1{
		size := len(res)
		for j := 0; j < size; j = j + 1{
			t := append(res[j], input[i])
			res = append(res, t)
		}
	}
	return res
}