/*
Remove the minimum number of invalid parentheses in order to make the input string valid. Return all possible results.

Note: The input string may contain letters other than the parentheses ( and ).

Example 1:

Input: "()())()"
Output: ["()()()", "(())()"]
Example 2:

Input: "(a)())()"
Output: ["(a)()()", "(a())()"]
Example 3:

Input: ")("
Output: [""]
*/
package main

import "fmt"

type Memory struct{
	list []string
}

func main(){

	//input := "()())()"
	input := "(a)())()"
	//input := ")("

	mem := Memory{}

	mem.removeInvalidParenthesis(input, 0, 0, []byte{'(', ')'})

	fmt.Println(mem)
}
//Stack to keep track when parentheses are invalid, which will happen when stack becomes negative
func (m *Memory) removeInvalidParenthesis(input string, start int, lastRemoveIndex int, par []byte){

	for stack, i := 0, start; i < len(input); i += 1{
		//increment the stack if open bracket
		if input[i] == par[0] {
			stack += 1
		}
		//decrement the stack if close bracket
		if input[i] == par[1]{
			stack -= 1
		}
		//if stack is positive, which means parenthesis is valid, so continue with the next element
		//It will take in other elements which are not brackets
		if stack >= 0{
			continue
		}
		//Else we need to remove the bracket and do recursive operation
		/*
		To make the prefix valid, we need to remove a ‘)’. The problem is: which one? The answer is any one in the prefix.
		However, if we remove any one, we will generate duplicate results, for example: s = ()), we can remove s[1] or s[2]
		but the result is the same (). Thus, we restrict oneself to remove the first ) in a series of consecutive )s.

		*/
		for j := lastRemoveIndex; j <= i; j += 1{
			if input[j] == par[1] && (j == lastRemoveIndex || input[j-1] != par[1]){
				/*
				After the removal, the prefix is then valid. We then call the function recursively to solve the rest of the string.
				However, we need to keep another information: the last removal position. If we do not have this position,
				we will generate duplicate by removing two ‘)’ in two steps only with a different order.
				For this, we keep tracking the last removal position and only remove ‘)’ after that.
				*/
				m.removeInvalidParenthesis(input[:j] + input[j+1:], i, j, par)
			}
		}
		/*
		Don't underestimate this return. It's very important
		if inside the outer loop, it reaches the above inner loop. You have scanned the str_to_check up to count_i
		In the above inner loop, when construct the new_str_to_check, we include the rest chars after count_i
		and call remove with it.
		So after the above inner loop finishes, we shouldn't allow the outer loop continue to next round because self.remove in the
		inner loop has taken care of the rest chars after count_i
		*/
		return
	}

	/*
	Now one may ask. What about ‘(‘? What if s = ‘(()(()’ in which we need to remove ‘(‘?
	The answer is: do the same from right to left.
	However, a cleverer idea is: reverse the string and reuse the code!
	*/
	input = reverse(input)

	if par[0] == '('{
		m.removeInvalidParenthesis(input, 0, 0, []byte{')', '('})
	} else {
		m.list = append(m.list, input)
	}
}

func reverse(str string)string{

	runes := []rune(str)

	for i, j := 0, len(runes) -1; i < j; i, j = i+ 1, j - 1{
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}