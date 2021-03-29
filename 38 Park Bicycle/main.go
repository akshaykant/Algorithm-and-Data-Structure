//Amex
/*
You want to park your bicycle in a bicycle parking area where bike racks are aligned in a row.
There are already N bikes parked there(each bike is attached to exactly one rack,
but a rack can have multiple bikes attached to it). We call racks that already have bikes attached used.
You want to park your bike in a rack in the parking area according to the following criteria:
        the chosen rack must lie between te first and the last used racks(inclusive);
        the distance between the chose rack and any other used rack is as big as possible.
A description of the bikes already parked in the racks is given in a non-empty zero-indexed
array A: element A[K] denotes the position of the rack to which bike number K is attached (for 0 <= K < N).
The central position in the parking area is position 0. A positive value means that the rack is located A[K]
meters to the right of the central position 0; a negative value means that it's located |A[K]| meters to the left.
That, given a non-empty zero-indexed array A of N integers, returns the largest possible distance in meters
between the chosen rack and any other used rack.
Assume that:
        N is an integer within the range [2..100,000];
        each element of array A is an integer within the range [-1,000,000,000..1,000,000,000].
Complexity:
        expected worst-case time complexity isO(N*log(N));
        expected worst-case space complexity is O(N), beyond input storage (not counting the
		storage required for input arguments).
*/
package main

import (
	"fmt"
	"sort"
)
import "math"

func main() {
	fmt.Print(Solution([]int{8, 0, 3, 4, 2}))
}

//find the maximum distance between the two points, and then park the car in the middle of the two points.
func Solution(A []int) int {

	sort.Ints(A)
	res := math.MinInt32
	if len(A) == 2 {
		return (A[1] - A[0]) / 2
	}

	for i := 0; i < len(A)-1; i++ {

		//fmt.Println(i, A[i+1], A[i])
		if A[i+1]-A[i] > 1 {
			res = max(res, A[i+1]-A[i])
		}
	}

	return res / 2

}

func max(i, j int) int {

	if i > j {
		return i
	}
	return j
}
