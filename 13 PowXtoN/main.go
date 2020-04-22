/*
Implement pow(x, n), which calculates x raised to the power n (xn).

Example 1:

Input: 2.00000, 10
Output: 1024.00000
Example 2:

Input: 2.10000, 3
Output: 9.26100
Example 3:

Input: 2.00000, -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
Note:

-100.0 < x < 100.0
n is a 32-bit signed integer, within the range [−231, 231 − 1]
*/
package main

import "fmt"


func main(){
	x, n := 2.0000, -2.0

	res := pow(x, n)
	fmt.Println(res)


	res2 := pow2(x, n)
	fmt.Println(res2)
}

//Run-time : O(n), Space : O(1)
func pow(x float64, n float64)(result float64){

	if n == 0{
		return 1
	}
	if n == 1{
		return x
	}
	if n == -1{
		return 1/x
	}

	if n < 0{
		result = 1/x * pow(x, n + 1)
		return result
	}
	result = x * pow(x, n - 1)

	return result
}

//Another approach is Divide and Conquer
//Run-time : O(log n)
func pow2 (x float64, n float64)(result float64){

	//base
	if n == 0{
		return 1
	}
	if n == 1{
		return x
	}
	if n == -1{
		return 1/x
	}

	if int(n) % 2 == 0{
		if n < 0 {
			result = pow2(x, n/2) * pow2(x, n/2)
		} else {
			result = pow2(x, n/2) * pow2(x, n/2)
		}
	} else {
		if n < 0 {
			result = 1/x * pow2(x, n/2) * pow2(x, n/2)
		} else {
			result = x * pow2(x, n/2) * pow2(x, n/2)
		}
	}
	return result
}