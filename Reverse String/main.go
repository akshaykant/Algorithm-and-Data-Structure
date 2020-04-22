/*Reverse a string*/
/*Rune : Rune literals are 32-bit integer values (however they're untyped constants, so their type can change).
They represent unicode code-points. For example, the rune literal 'a' is actually the number 97.
For preserving Unicode combining characters such as "as⃝df̅" with Reverse String result "f̅ds⃝a", we need to convert
into rune and do the computation for string*/

package main

import "fmt"

func main(){
	var str string
	str = "The quick brown 狐 jumped over the lazy 犬 The quick brown 狐 jumped over the lazy 犬"

	output1, _ := reverseString1(str)
	fmt.Println("Better Solution : Rune", output1)

	output2, _ := reverseString2(str)
	fmt.Println("Range:",output2)
}

//Approach 1
//convert string into rune literal
func reverseString1(input string) (output string, error string) {

	if input == ""{
		return "", "ERROR : empty string"
	}


	runes := []rune(input)

	/*Two pointer solution : Opposite-directional. One at the start and other at the end, they move close
	to each other and meet in the middle (-> <-)*/
	for i,j := 0,len(runes)-1; i < j; i,j = i+1, j-1 {

		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes), ""
}

//Approach 2
//Strings are immutable and hence this is inefficient solution.
func reverseString2 (input string) (output string, error string){
	if input == ""{
		return "", "ERROR : empty string"
	}
	rev := ""
	for _ ,v := range input{
		rev = string(v) + rev
	}
	return rev, ""
}