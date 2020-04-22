/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
*/
package main

import "fmt"

type Stack struct{
	m []byte
	top int
}
func main(){
	//str := "()"
	//str := "()[]{}"
	//str := "([)]"
	//str := "(]"
	str := "{[]}"

	m := make([]byte, len(str))
	top := -1
	stack := &Stack{
		m,
		top,

	}
	result := stack.isParentheses(str)

	fmt.Println(result)
}

func (stack *Stack) isParentheses(input string)(result bool){

	if input == "" {
		return false
	}
	//Odd number of elements means missing brackets
	if len(input) % 2 != 0 {
		return false
	}
	//Iterate through the string
	//Push to Stack, if is open bracket.Pop to Stack if it is close bracket.
	for  i:= 0; i < len(input); i = i + 1{
		if input[i] == '(' || input[i] == '{' || input[i] == '['{
			//Push
			stack.push(input[i])
		}
		if input[i] == ')' || input[i] == '}' || input[i] == ']'{
			//Pop
			b, err := stack.pop()
			if err != ""{
				return false
			}
			var bCompare byte

			if input[i] == ')' {
				bCompare = '('
			}
			if input[i] == '}'{
				bCompare = '{'
			}
			if input[i] == ']'{
				bCompare = '['
			}
			if bCompare != b{
				return false
			}
		}
	}
	return true
}
func (stack *Stack) push (ele byte){

	stack.top += 1
	stack.m[stack.top] = ele

}

func (stack *Stack) pop () (b byte, error string){
	if stack.top < 0 {
		return b, "Error"
	}
	b = stack.m[stack.top]
	//Set stack to nil, so that garbage collector can clean the memory
	stack.m[stack.top] = 0
	stack.top -= 1
	return b,""
}
