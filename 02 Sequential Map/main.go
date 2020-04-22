/*Iterate through a map keys in a sorted way
Map does not iterate in a sequence way. It gives out random values
*/
package main

import (
	"fmt"
)

func main(){

	m := make(map[int]int)
	for i := 0; i <= 20; i = i + 1{
		m[i] = i
	}
	result, _ := iterateMap(m)

	fmt.Println(result)

	for _, v := range result {
		fmt.Println(v, m[v])
	}
}

func iterateMap(input map[int]int)(result []int, err string){

	if input == nil{
		return []int{}, "Error: empty input"
	}

	//Iterate through the map and store each key in a slice
	for k,_ := range input {
		result = append(result, k)
	}

	//Sort the slice
	//sort.Ints(result)
	result, _ = sort(result)

	return result, ""
}

func sort(in []int)(re []int, err string){

	for i := 0; i < len(in) -1 ; i = i + 1{
		for j := 0; j < len(in) - i - 1; j = j + 1{
			if in[j] > in[j + 1]{
				//swap
				in[j], in[j + 1] = in[j + 1], in[j]
			}
		}
	}
	return in, ""
}