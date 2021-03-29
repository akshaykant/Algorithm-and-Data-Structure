/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
 "((()))",
 "(()())",
 "(())()",
 "()(())",
 "()()()"
]

*/

package main

import "fmt"

type Memory struct {
	list []string
}

func main() {
	mem := Memory{}

	n := 3

	mem.generateBracket(n, 0, 0, []byte{})

	fmt.Println(mem.list)

}

/*
let's only add them when we know it will remain a valid sequence.
We can do this by keeping track of the number of opening and closing brackets we have placed so far.

We can start an opening bracket if we still have one (of n) left to place.
And we can start a closing bracket if it would not exceed the number of opening brackets.
*/
func (m *Memory) generateBracket(n int, open int, close int, cur []byte) {

	//base condition
	if len(cur) == n*2 {
		m.list = append(m.list, string(cur))
	}

	if open < n {
		m.generateBracket(n, open+1, close, append(cur, '('))
	}
	if close < open {
		m.generateBracket(n, open, close+1, append(cur, ')'))
	}

}
