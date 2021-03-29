package main

import "fmt"

/*
 * Complete the 'selectStock' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER saving
 *  2. INTEGER_ARRAY currentValue
 *  3. INTEGER_ARRAY futureValue
 */

func selectStock(saving int32, currentValue []int32, futureValue []int32) int32 {
	// Write your code here
	selectedValue := make([]int32, len(currentValue))
	//loop through each value and check if it is profit or loss. Incase of loss add 0 to handle worst combination calculation
	for i := range currentValue {
		val := currentValue[i] - futureValue[i]
		if val < 0 {
			val = 0
		}
		selectedValue[i] = val
	}
	return knapsack(selectedValue, currentValue, saving)
}

func knapsack(selectedValues []int32, cuurentValues []int32, savings int32) int32 {
	sack := make([][]int32, savings+1)

	for v := range sack {
		sack[v] = make([]int32, len(cuurentValues)+1)
	}

	var i int
	for i = 1; i < (len(cuurentValues)); i++ {
		var j int32
		for j = 1; j < savings; j++ {

			if i == 0 || j == 0 {
				sack[i][j] = 0
			} else if cuurentValues[i-1] <= j {
				sack[i][j] = max(selectedValues[i-1]+sack[i-1][j-cuurentValues[i-1]], sack[i-1][j])
			} else {
				sack[i][j] = sack[i-2][j-1]
			}
		}
	}
	return sack[len(cuurentValues)][savings]
}

func max(a int32, b int32) (res int32) {
	if a > b {
		res = a
	}
	if b > a {
		res = b
	}
	if a == b {
		res = a
	}
	return res
}

func main() {
	var saving int32 = 250
	currentValue := []int32{175, 133, 109, 201, 97}
	futureValue := []int32{200, 125, 128, 228, 133}

	res := selectStock(saving, currentValue, futureValue)
	fmt.Println(res)
}
