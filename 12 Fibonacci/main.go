/*
The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence,
such that each number is the sum of the two preceding ones, starting from 0 and 1.
That is,

F(0) = 0,  F(1) = 1
F(N) = F(N - 1) + F(N - 2), for N > 1.
Given N, calculate F(N).



Example 1:

Input: 2
Output: 1
Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
Example 2:

Input: 3
Output: 2
Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
Example 3:

Input: 4
Output: 3
Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.


Note:

0 ≤ N ≤ 30.
*/
package main

import "fmt"

//Used in 2nd function
type memory struct {
	mem map[int]int
}

func main() {

	result1 := calculateFibonacci1(6)

	fmt.Println(result1)

	//map to store the repeated values
	m := memory{
		make(map[int]int),
	}

	result2 := m.calculateFibonacci2(6)

	fmt.Println(result2)

	result3 := calculateFibonacci3(6)

	fmt.Println(result3)

	result4 := calculateFibonacci4(6)

	fmt.Println(result4)

}

//Recursion
//Time- O(2^n) ; Space - O(2^n) for all stack calls
func calculateFibonacci1(n int) (result int) {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	result = calculateFibonacci1(n-1) + calculateFibonacci1(n-2)

	return result

}

//DP = recursion + memoization
//Time - O(n); Space O(n)
//take the number as input to compute and the memory as struct to hold all the computed numbers.
// If number if not present in memory,return the value else calculate and add to the memory
func (m *memory) calculateFibonacci2(n int) (result int) {

	//Check if memory holds the value
	val, boolean := m.mem[n]
	if boolean {
		return val
	}
	//base condition
	if n <= 1 {
		m.mem[n] = n
		return n
	}
	result = m.calculateFibonacci2(n-1) + m.calculateFibonacci2(n-2)
	m.mem[n] = result
	return result

}

//DP Bottom Up approach (Optimized runtime)
//Time - O(n) , Space - O(n)
//Store all possible values in memory and result in the value. This memory can later be used to get the value in Constant Time
func calculateFibonacci3(n int) (result int) {

	m := make(map[int]int)

	//Store all the possible values in bottom up manner
	for i := 0; i <= n; i = i + 1 {
		if i <= 1 {
			m[i] = i
		} else {
			m[i] = m[i-1] + m[i-2]
		}
	}

	result, _ = m[n]

	return result
}

//DP Bottom Up approach (Optimized space)
//Time - O(N), Space - O(1)
func calculateFibonacci4(n int) (result int) {

	if n == 0 {
		return 0
	}

	first, second := 0, 1

	for i := 1; i <= n; i = i + 1 {
		if i == 1 {
			second = 1
		} else {
			first = first + second
			first, second = second, first
		}
	}
	result = second
	return result
}
