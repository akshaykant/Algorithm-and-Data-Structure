/*
Given a string, reduce the string by removing 3 or more consecutive identical characters. You should greedily remove characters from left to right.

Example 1:

Input: "aaabbbc"
Output: "c"
Explanation:
1. Remove 3 'a': "aaabbbc" => "bbbc"
2. Remove 3 'b': "bbbc" => "c"

Example 2:

Input: "aabbbacd"
Output: "cd"
Explanation:
1. Remove 3 'b': "aabbbacd" => "aaacd"
2. Remove 3 'a': "aaacd" => "cd"
Example 3:

Input: "aabbccddeeedcba"
Output: ""
Explanation:
1. Remove 3 'e': "aabbccddeeedcba" => "aabbccdddcba"
2. Remove 3 'd': "aabbccdddcba" => "aabbcccba"
3. Remove 3 'c': "aabbcccba" => "aabbba"
4. Remove 3 'b': "aabbba" => "aaa"
5. Remove 3 'a': "aaa" => ""
Example 4:

Input: "aaabbbacd"
Output: "acd"
Explanation:
1. Remove 3 'a': "aaabbbacd" => "bbbacd"
2. Remove 3 'b': "bbbacd" => "acd"

Follow-up:
What if you need to find the shortest string after removal?

Example:

Input: "aaabbbacd"
Output: "cd"
Explanation:
1. Remove 3 'b': "aaabbbacd" => "aaaacd"
2. Remove 4 'a': "aaaacd" => "cd"
*/
package main

import "fmt"

type Values struct {
	key   byte
	count int
}

type Stack struct {
	mem []Values
	top int
}

func main() {
	//in := "aaabbbc"
	//in := "aabbbacd"
	//in := "aabbccddeeedcba"
	in := "aaabbbacd"

	stack := &Stack{make([]Values, len(in)), -1}

	out := stack.candycrush(in)
	out2 := stack.candycrush_followup(in)

	fmt.Println(out)
	fmt.Println(out2)
}

/*
- Iterate through each character in the string.
- Peek the Stack, if stack is empty or element value is not equal to the character.
Pop the stack if count is greater than 3 and Push the new character to the Stack.
- If current character is same as the top in the stack, Pop, increment the count and Push.
*/
func (stack *Stack) candycrush(str string) string {

	i := 0
	for i < len(str) {
		v, err := stack.Peek()

		//If Stack is empty
		if err != "" {
			stack.Push(Values{str[i], 1})
			i = i + 1
		} else if v.key != str[i] {
			if v.count >= 3 {
				stack.Pop()
				continue //Need not to increment the character as it can shrink the string once added.
			}
			stack.Push(Values{str[i], 1})
			i = i + 1
		} else if v.key == str[i] {
			//Pop the stack to update the counter
			stack.Pop()
			stack.Push(Values{v.key, v.count + 1})
			//edge condition : when reached the end of the string and value was same
			if i == len(str)-1 {
				v, _ := stack.Peek()
				if v.count >= 3 {
					stack.Pop()
				}
			}
			i = i + 1
		}
	}

	out := make([]byte, 0)
	//Pop out each item and add it to the list to form the final string
	for !stack.isEmpty() {
		val, _ := stack.Pop()
		for i := 0; i < val.count; i = i + 1 {
			out = append(out, val.key)
		}
	}

	return reverse(out)
}

func (stack *Stack) candycrush_followup(str string) string {

	i := len(str) - 1
	for i >= 0 {
		v, err := stack.Peek()

		//If Stack is empty
		if err != "" {
			stack.Push(Values{str[i], 1})
			i = i - 1
		} else if v.key != str[i] {
			if v.count >= 3 {
				stack.Pop()
				continue //Need not to increment the character as it can shrink the string once added.
			}
			stack.Push(Values{str[i], 1})
			i = i - 1
		} else if v.key == str[i] {
			//Pop the stack to update the counter
			stack.Pop()
			stack.Push(Values{v.key, v.count + 1})
			//edge condition : when reached the end of the string and value was same
			if i == 0 {
				v, _ := stack.Peek()
				if v.count >= 3 {
					stack.Pop()
				}
			}
			i = i - 1
		}
	}

	out := make([]byte, 0)
	//Pop out each item and add it to the list to form the final string
	for !stack.isEmpty() {
		val, _ := stack.Pop()
		for i := 0; i < val.count; i = i + 1 {
			out = append(out, val.key)
		}
	}

	return reverse(out)
}

func reverse(in []byte) string {
	for i, j := 0, len(in)-1; i < j; i = i + 1 {
		in[i], in[j] = in[j], in[i]
	}
	return string(in)
}
func (stack *Stack) isEmpty() bool {

	if stack.top < 0 {
		return true
	}
	return false
}

func (stack *Stack) Push(v Values) {
	stack.top = stack.top + 1
	stack.mem[stack.top] = v
}

func (stack *Stack) Pop() (Values, string) {
	if stack.isEmpty() {
		return Values{}, "Error : Empty Stack"
	}
	v := stack.mem[stack.top]
	stack.mem[stack.top] = Values{}
	stack.top = stack.top - 1

	return v, ""
}

func (stack *Stack) Peek() (Values, string) {
	if stack.isEmpty() {
		return Values{}, "Error : Empty String"
	}
	return stack.mem[stack.top], ""
}
