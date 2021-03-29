//Microsoft

/*
In a given string some of the characters are replaced by question mark,
and you can replace question mark with any character.
Given such a string find total number of palindrome that can created.
String contains only [a-z] characters and question marks can also be only replaced by [a-z].

Example:
Input String: String str=”a??a”
Output: "azza" or "aaaa"

*/

package main

import (
	"fmt"
)

func Solution(S string) string {
	// write your code in Go 1.4
	if S == "" || len(S) == 0 {
		return "NO"
	}
	runes := []rune(S)

	for first, last := 0, len(runes)-1; first <= last; first, last = first+1, last-1 {
		if runes[first] == '?' && runes[last] == '?' {
			runes[first] = 'a'
			runes[last] = 'a'
		} else if runes[first] == '?' {
			runes[first] = runes[last]
		} else if runes[last] == '?' {
			runes[last] = runes[first]
		} else if runes[first] != runes[last] {
			return "NO"
		}
	}
	return string(runes)
}

func main() {
	fmt.Println(Solution("?ab??a"))
	fmt.Println(Solution("a?a"))
	fmt.Println(Solution("???"))
}
