/*
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true
Example 2:

Input: "race a car"
Output: false
*/
package main

import (
	"fmt"
)

func main() {
	str := "A man, a plan, a canal: Panama"

	result := isPalindrome(str)

	fmt.Println(result)
}

func isPalindrome(input string) (result bool) {

	if input == "" {
		return false
	}

	//two pointers for checking each rune, one from start and other from end

	for first, last := 0, len(input)-1; first < last; {

		//checkForDigitAndAlphabet
		isFirstAlphanumeric := checkForDigitAndAlphabet(input[first])
		isLastAlphanumeric := checkForDigitAndAlphabet(input[last])

		//Pointer increase
		if first < last && !isFirstAlphanumeric {
			first = first + 1
		}
		//Pointer Decrease
		if first < last && !isLastAlphanumeric {
			last = last - 1
		}

		if isFirstAlphanumeric && isLastAlphanumeric {
			//convert to lowercase and compare
			if toLowerCase(input[first]) != toLowerCase(input[last]) {
				return false
			}

			first, last = first+1, last-1

		}

	}

	return true
}

func checkForDigitAndAlphabet(b byte) (res bool) {

	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
		return true
	}

	return false
}

func toLowerCase(b byte) (res byte) {
	if b >= 'A' && b <= 'Z' {
		res = b + 'a' - 'A'
	} else {
		res = b
	}

	return res
}
