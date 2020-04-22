/*
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.
*/
package main

import "fmt"

func main(){
	//input := []int{0,1,0,3,12}
	input := []int{0,0,0,0,0,1,0,3,12.0,0,1,0,3,12}

	out, err := moveZero(input)

	if err != ""{
		fmt.Println(err)
	}

	if err == ""{
		fmt.Println(out)
	}
}

/*
In-place means we should not be allocating any space for extra array.
But we are allowed to modify the existing array. However, as a first step,
try coming up with a solution that makes use of additional space. For this
problem as well, first apply the idea discussed using an additional array
and the in-place solution will pop up eventually.
*/

/*
A two-pointer approach could be helpful here.
The idea would be to have one pointer for iterating the array and another pointer
that just works on the non-zero elements of the array.
*/

//RunTime : O(N) - N - length of array
//SpaceTime : O(1)
func moveZero(input []int)([]int, string){

	if input == nil{
		return []int{}, "ERROR: empty array"
	}

	nonZeroPtr := 0

	for i := 0; i < len(input); i += 1{

		if input[i] != 0{
			input[nonZeroPtr] = input[i]
			nonZeroPtr += 1
		}
	}

	for i := nonZeroPtr; i < len(input); i += 1{
		input[i] = 0
	}


	return input, ""
}