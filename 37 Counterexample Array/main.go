//Amex
/*
The following find_min function should return the smallest integer from a given array A.

int findmin (int [] A) {
	int ans = 0;
	for (int i = 1; i < A. length; i++) {
		if (ans > A[i] ){
			ans = A[i];
		}
	}
	return ans;
}

Unfortunately it is an incorrect implementation. In other words, when the function is called with certain parameters,
it returns the wrong answer. Your task is to generate a counterexample, i.e.
an array A consisting of N integers such that the find_min function returns the wrong answer.
Write a function:
class Solution { public int[] solution(int N); }

that, given an integer N, returns an array A consisting of N integers which describes a counterexample.
Example: Given N = 4, your function may return [4, 2,4, 5].
It is a counterexample, because calling the find_min function with this array returns 0, but the correct answer is 2.
Your function may also return another counterexample;
for example, [10, 567, 99,456].
Assume that
• N is an integer within the range 11..1,0001.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Print(Solution(4))
}

// Generate an random int array of size N which has all elements > 0

func Solution(N int) []int {
	// write your code in Go 1.4
	//Use better seed for better randomness.
	rand.Seed(time.Now().UnixNano())
	res := make([]int, N)

	for i := 0; i < N; i++ {
		r := rand.Intn(i + 1)
		res[i] = r + 1
	}
	return res
}
