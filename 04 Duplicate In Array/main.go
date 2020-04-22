/*
Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements that appear twice in this array.

Could you do it without extra space and in O(n) runtime?

Example:
Input:
[4,3,2,7,8,2,3,1]

Output:
[2,3]
*/
package main

import (
	"fmt"
	"math"
)

func main(){
	list := []int{4,3,2,7,8,2,3,1}
	fmt.Println(list)
	list = findDuplicate(list)
	fmt.Println(list)
}

func findDuplicate(input []int)(result []int){

	//var index int

	for i := 0; i < len(input); i = i + 1{
		//check for the element - 1 as index of array. This is because of the assumption that length of array is biggest element in the array.
		// So array index will be 0 to n-1. Check if the element is positive, change to negative.


		index := int(math.Abs(float64(input[i]))) - 1

		// Element at array index i subtract 1, is the new index
		if input[index] > 0 { //Mark negative only if element at the index in positive, else it is already visited
			input[index] = - input[index]
		} else {
			result = append(result, index + 1)
		}
	}

	return result
}