/*
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
Example:

Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
*/
package main

import "fmt"

func main() {

	num1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	num2 := []int{2, 5, 6}
	n := 3


	out := merge2array(num1, m, num2, n)

	fmt.Println(out)
}

//Two pointers running from end of each list. And another pointer end of length of first list, to preserve memory.
//Compare both elements and decrement the pointer until reaches the starting of array
func merge2array(list1 []int, m int, list2 []int, n int)[]int{

	//two pointers at the end of list element
	ptr1, ptr2 := m-1, n-1

	//ptr to the end of list memory
	ptrList := m + n - 1

	for ptr1 >= 0 && ptr2 >= 0 {

		if list1[ptr1] < list2[ptr2]{
			list1[ptrList] = list2[ptr2]
			ptr2 -= 1
		} else {
			list1[ptrList] = list1[ptr1]
			ptr1 -= 1
		}
		ptrList -= 1
	}

	//If list2 contains any other element because they are smaller than list2, add them to the starting of list 1

	copy(list1[0:ptr2+1] , list2[0:ptr2+1])

	return list1
}