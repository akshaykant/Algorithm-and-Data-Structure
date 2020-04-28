/*
https://www.youtube.com/watch?v=aL6cU5dWmWM&feature=emb_logo

https://www.youtube.com/watch?v=vZtCKL_OwdA

https://leetcode.com/problems/partition-equal-subset-sum/discuss/90592/01-knapsack-detailed-explanation

Given a non-empty array containing only positive integers, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

Note:

Each of the array element will not exceed 100.
The array size will not exceed 200.


Example 1:

Input: [1, 5, 11, 5]

Output: true

Explanation: The array can be partitioned as [1, 5, 5] and [11].


Example 2:

Input: [1, 2, 3, 5]

Output: false

Explanation: The array cannot be partitioned into equal sum subsets.
*/
package main

import "fmt"

func main(){

	in := []int{1,5,11,5}
	// It asks us to find if there is a subset whose sum equals to target value. For this problem, the target value is exactly the half of sum of array.
	sum := 0
	for _, v := range in{
		sum += v
	}

	target := sum / 2
	res := sumSubArray(in, sum, target)

	fmt.Println(res)
}

func sumSubArray(input []int, sum int, target int) bool{
	//If sum is odd, it cannot be divided into equal subsets
	//if the array sum is not even, we cannot partition it into 2 equal subsets
	if sum % 2 == 1{
		return false
	}
	//Create a 2D array to store the result for the DP. For various length of the input and different target
	subset := make([][]bool, len(input))
	for i := range subset {
		subset[i] = make([]bool, target + 1)
	}

	//deal with the first row
	if input[0] <= target{
		subset[0][input[0]] = true
	}

	//deal with first column
	for i := 0; i < len(input); i += 1{
		subset[i][0] = true
	}

	//deal with rest of the array
	for i := 1; i < len(subset); i += 1{
		for j:= 1; j < len(subset[0]); j += 1{
			if j < input[i]{
				subset[i][j] = subset[i-1][j]
			} else {
				subset[i][j] = subset[i-1][j] || subset[i-1][j-input[i]]
			}
		}
	}

	return subset[len(subset) - 1][len(subset[0]) - 1]
}