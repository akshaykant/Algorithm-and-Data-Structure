//Microsoft
/*
Given a string s and an array of integers cost where cost[i] is the cost of deleting the ith character in s.

Return the minimum cost of deletions such that there are no two identical letters next to each other.

Notice that you will delete the chosen characters at the same time, in other words, after deleting a character,
the costs of deleting other characters will not change.

Example 1:

Input: s = "abaac", cost = [1,2,3,4,5]
Output: 3
Explanation: Delete the letter "a" with cost 3 to get "abac" (String without two identical letters next to each other).
Example 2:

Input: s = "abc", cost = [1,2,3]
Output: 0
Explanation: You don't need to delete any character because there are no identical letters next to each other.
Example 3:

Input: s = "aabaa", cost = [1,2,3,4,1]
Output: 2
Explanation: Delete the first and the last character, getting the string ("aba").


Constraints:

s.length == cost.length
1 <= s.length, cost.length <= 10^5
1 <= cost[i] <= 10^4
s contains only lowercase English letters.
*/
package main

import "fmt"

//Maintain the running sum and max value for repeated letters.
func Solution(S string, C []int) int {

	if S == "" {
		return 0
	}

	if len(C) == 0 {
		return 0
	}
	// write your code in Go 1.4
	runes := []rune(S)
	var totalCost int

	i := 1

	for i < len(runes) {
		if runes[i] == runes[i-1] {
			cost1 := C[i-1]
			cost2 := C[i-1]
			for i < len(runes) && (runes[i] == runes[i-1]) {
				cost1 += C[i]
				cost2 = max(cost2, C[i])
				i++
			}
			totalCost += cost1
			totalCost -= cost2
		} else {
			i++
		}
	}
	return totalCost

}

func max(i, j int) int {

	if i > j {
		return i
	}
	return j
}

func main() {

	fmt.Print(Solution("aabaa", []int{1, 2, 3, 4, 1}))
}
