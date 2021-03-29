/*
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, add spaces in s to construct a sentence where each word is a valid dictionary word. Return all such possible sentences.

Note:

The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.
Example 1:

Input:
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
Output:
[
 "cats and dog",
 "cat sand dog"
]
Example 2:

Input:
s = "pineapplepenapple"
wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
Output:
[
 "pine apple pen apple",
 "pineapple pen apple",
 "pine applepen apple"
]
Explanation: Note that you are allowed to reuse a dictionary word.
Example 3:

Input:
s = "catsandog"
wordDict = ["cats", "dog", "sand", "and", "cat"]
Output:
[]
*/

package main

import "fmt"

func main() {

	s := "pineapplepenapple"
	wordDict := []string{"apple", "pen", "applepen", "pine", "pineapple"}

	/*s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}*/

	/*s := "catsandog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}*/

	list := wordBreak(s, wordDict)

	fmt.Println(list)

}

/*
Insert all the words in the HashMap.
It is solved by traversing through the entire string - DFS,
DP + Backtracking when we find the word and all subsequent words.
Backtrack to find other combinations.
Memoization so that we do not traverse again for the already seen path. We are doing this by storing the possible combination for a particular position.
*/
func wordBreak(str string, wordDict []string) []string {

	wordMap := make(map[string]bool)

	for _, word := range wordDict {
		wordMap[word] = true
	}

	res := dfs(str, wordMap, make(map[int][]string), 0)

	return res
}

func dfs(str string, wordMap map[string]bool, record map[int][]string, pos int) []string {

	//memoization : check if we already have traversed for a particular position and hold all the possible list for that position
	if result, ok := record[pos]; ok {
		return result
	}

	//list to hold the string from word dictionary
	var result []string

	//Iterating till the length of the string as substring works till len-1
	for i := pos + 1; i <= len(str); i += 1 {
		substring := str[pos:i]

		//If word is present in the Hash Map, recursively, find all other possible list from the updated position
		if wordMap[substring] {

			if i != len(str) {
				rest := dfs(str, wordMap, record, i)

				//append the matched substring along with the rest to the result
				//result = append(result, substring)
				for _, v := range rest {
					result = append(result, substring+" "+v)
				}
			} else {
				//if end of the string is reached, no need to recurse
				result = append(result, substring+"\n")

				break
			}
		}
	}

	//Memoization : update the seen list for the postion
	record[pos] = result

	return result
}
