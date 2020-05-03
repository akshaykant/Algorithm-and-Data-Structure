/*
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:

Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
Example 2:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/
package main

import "fmt"

func main(){

	//in := []int{7,1,5,3,6,4}
	in := []int{7,6,4,3,1}

	res := buysell(in)

	fmt.Println(res)
}

/*
Time complexity : O(n)O(n). Only a single pass is needed.

Space complexity : O(1)O(1). Only two variables are used.

The points of interest are the peaks and valleys in the given graph. We need to find the largest peak following the smallest valley. 
We can maintain two variables - minprice and maxprofit corresponding to the smallest valley and maximum profit
(maximum difference between selling price and minprice) obtained so far respectively.
*/
func buysell(in []int) int{

	//Atleast 2 days of data is needed for profit calculation
	if len(in) < 2{
		return 0
	}

	/*
	We'll greedily update min_price and max_profit, so we initialize
	them to the first price and the first possible profit
	*/
	minPrice := in[0]
	maxProfit := in[1] - in[0]

	/*
	Start at the second (index 1) time. We can't sell at the first time, since we must buy first,
	and we can't buy and sell at the same time! If we started at index 0, we'd try to buy *and* sell at time 0.
	This would give a profit of 0, which is a problem if our max_profit is supposed to be *negative*--we'd return 0.
	*/
	for currentPrice := 1; currentPrice < len(in); currentPrice++{


		/*
		See what our profit would be if we bought at the min price and sold at the current price.
		And Update max_profit if we can do better
		*/
		maxProfit = max(maxProfit, in[currentPrice] - minPrice)

		//Update min_price so it's always the lowest price we've seen so far
		minPrice = min(minPrice, in[currentPrice])
	}

	if maxProfit < 0 {
		return 0
	}
	
	return maxProfit
}

func min(i, j int)int{

	if i < j {
		return i
	}
	return j
}

func max(i, j int)int{

	if i > j {
		return i
	}
	return j
}