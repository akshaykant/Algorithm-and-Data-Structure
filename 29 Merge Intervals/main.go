/*
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
*/
package main

import "fmt"
import "sort"

func main() {

	/*input := [][]int{{2,6},{15, 18},{8,10},{1, 3}}*/
	input := [][]int{{1, 4}, {4, 5}}

	res := mergeInterval(input)

	fmt.Println(res)
}

/*
O(nlgn) time and O(n)O(n) space.
If we sort the intervals by their start value, then each set of intervals that can be merged will appear as a contiguous "run" in the sorted list.
*/
func mergeInterval(intervals [][]int) [][]int {

	//Sort the given slice
	sort.Slice(intervals[:], func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	result := make([][]int, 0)

	//Iterate over the intervals and compare for [i[j] with the last element in result
	// Of two intervals, if the [j]] of result(first) is less than equal to the [i] of interval(second), update the result or append it
	for i, interval := range intervals {
		//For first element, insert in the result
		if i == 0 {
			result = append(result, interval)
			continue
		}

		//get the last added element in the result,as this need to be merged
		res := result[len(result)-1]

		if interval[0] <= res[1] && res[1] < interval[1] {
			result[len(result)-1] = []int{res[0], interval[1]} //update the last added to result with the start range of result and end range of interval
		}
		if res[1] < interval[0] {
			result = append(result, interval)
		}
	}

	return result
}
