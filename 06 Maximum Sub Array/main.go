/*Maximum Subarray
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.*/

package main

import "fmt"

func main(){
	arr := []int{-2,1,-3,4,-1,2,1,-5,4}

	sum, _ := calculateSum(arr)

	fmt.Println(sum)

}

func calculateSum(input []int)(sum int, error string) {
	if input == nil {
		return 0, "ERROR: empty array"
	}
	//Assign local sum i.e. sum of the sub array of previous index and global sum i.e. the sum of sub array on the global level.
	sumLocal, sumGlobal := input[0], input[0]

	for i := 1; i < len(input); i = i + 1 {
		//calculate local maximum sum
		sumLocal = max(input[i], input[i]+sumLocal)

		//calculate global maximum sum
		sumGlobal = max(sumLocal, sumGlobal)
	}
	sum = sumGlobal

	return sum, ""
}

func max(a int, b int)(res int){

	if a > b {
		res = a
	}

	if b > a {
		res = b
	}

	if a == b {
		res = a
	}
	return res
}
