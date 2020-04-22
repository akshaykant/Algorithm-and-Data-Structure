/*
In an alien language, surprisingly they also use English lowercase letters, but possibly in a different order.
The order of the alphabet is some permutation of lowercase letters.

Given a sequence of words written in the alien language, and the order of the alphabet, return true if and only
if the given words are sorted lexicographically in this alien language.



Example 1:

Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
Output: true
Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.
Example 2:

Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
Output: false
Explanation: As 'd' comes after 'l' in this language, then words[0] > words[1], hence the sequence is unsorted.
Example 3:

Input: words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
Output: false
Explanation: The first three characters "app" match and the second string is shorter (in size.)
According to lexicographical rules "apple" > "app", because 'l' > '∅', where '∅' is defined as the blank
character which is less than any other character (More info).


Constraints:

1 <= words.length <= 100
1 <= words[i].length <= 20
order.length == 26
All characters in words[i] and order are English lowercase letters.
*/
package main

import "fmt"

func main(){

	words := []string{"hello","leetcode"}
	order := "hlabcdefgijkmnopqrstuvwxyz"

/*	words := []string{"word","world","row"}
	order := "worldabcefghijkmnpqstuvxyz"*/

/*	words := []string{"apple","app"}
	order := "abcdefghijklmnopqrstuvwxyz"*/

	bool := isAlienDictionary(words, order)

	fmt.Println(bool)
}
/*The words are sorted lexicographically if and only if adjacent words are.
This is because order is transitive: a <= b and b <= c implies a <= c.*/
//Time Complexity: O(C), where C is the total content of words.
//The outer loop runs for C times which is the content of words or the length of the array. The inner loop is constant.
//Space Complexity: O(1)
func isAlienDictionary(words []string, order string) bool {


	//create a int list to have the ASCII value of the Alien Dictionary
	listAlien := make([]int, 26)
	for i, v := range order {
		listAlien[v - 'a'] = i
	}

	//Iterate through the list of words and have two pointers to compare each word
	for i := 0; i < len(words) - 1; i += 1{
		word1, word2 := words[i], words[i+1]
		//We need to check till the min len of word, to avoid overflow
		minLen := minStrLen(word1, word2)
		//Used to check if array till an index are same
		var j int

		for j = 0; j < minLen; j += 1{

			//if each character in the word does not match
			if word1[j] != word2[j]{
				//character as per Alien dictionary are not in lexical order
				if listAlien[word1[j] - 'a'] > listAlien[word2[j] - 'a']{
					return false
				}
				break
			}
		}

		//If character are same till a length like "abc" & "abcd"
		if j == minLen && len(word1) > len(word2){
			return false
		}

	}
	return true
}

func minStrLen(word1 string, word2 string) int {
	if len(word1) > len(word2){
		return len(word2)
	}else
	if len(word2) > len(word1){
		return len(word1)
	}else
	if len(word1) == len(word2){
		return len(word2)
	}

	return 0
}