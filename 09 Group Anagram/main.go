/*
Given an array of strings, group anagrams together.

Example:

Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
 ["ate","eat","tea"],
 ["nat","tan"],
 ["bat"]
]
Note:

All inputs will be in lowercase.
The order of your output does not matter.
*/
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	group1, err1 := computeGroupAnagram1(input)

	group2, err2 := computeGroupAnagram2(input)

	if err1 == "" {
		for k, v := range group1 {
			fmt.Println(k, v)
		}
	}

	if err1 != "" {
		fmt.Println(err1)
	}

	if err2 == "" {
		for k, v := range group2 {
			fmt.Println(k, v)
		}
	}

	if err2 != "" {
		fmt.Println(err2)
	}

}

//Runtime : O(NKlogK) - N : number of words in the string. K is the max length the word
//Spacetime : O(NK) - hash map
func computeGroupAnagram1(input []string) (groupAnagram map[string][]string, err string) {

	if input == nil {
		return nil, "ERROR: empty input"
	}

	groupAnagram = make(map[string][]string)
	//Sort each word in the input list and add the map, to group
	for _, v := range input {
		sortWord := sortString(v)
		groupAnagram[sortWord] = append(groupAnagram[sortWord], v)

	}
	return groupAnagram, ""
}

func sortString(str string) string {

	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

//Runtime : O(NK) - N : number of words in the string. K is the max length of the word. Constant time to create the string
//Spacetime : O(NK) - hash map
func computeGroupAnagram2(input []string) (map[string][]string, string) {

	if input == nil {
		return nil, "ERROR: empty input"
	}

	groupAnagram := make(map[string][]string)

	//Count the number of character occurrence of the word, and use that as the key in the map.
	for _, word := range input {

		countLetters := make([]int, 26)
		for _, letter := range word {

			countLetters[letter-'a'] += 1

		}

		str := make([]string, 26)
		//Constant time loop - to create the key
		for i := 0; i < 26; i += 1 {
			str[i] = strconv.Itoa(countLetters[i])
		}
		key := strings.Join(str, "")

		groupAnagram[key] = append(groupAnagram[key], word)
	}
	return groupAnagram, ""
}
