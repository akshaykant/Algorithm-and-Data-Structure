//Microsoft
//https://leetcode.com/discuss/interview-question/574309/codility-plane-seat-reservation
//https://leetcode.com/problems/cinema-seat-allocation/
package main

import (
	"fmt"
	"strings"
)

func Solution(N int, S string) int {
	//if no seats are reserved, each row will have 2 set of 4 seats(family seat)
	if S == "" || len(S) == 0 {
		return N * 2
	}
	reservedSeatsMap := make(map[string][]string)

	seats := strings.Split(S, " ")

	//Mapping reserved seats at every row
	for _, seat := range seats {
		l := len(seat)
		fmt.Println("len", l)
		fmt.Println("seat", seat)
		fmt.Println("number", seat[:l-1])
		fmt.Println("letter", seat[l-1:])
		reservedSeatsMap[seat[:l-1]] = append(reservedSeatsMap[seat[:l-1]], seat[l-1:])
	}
	fmt.Println(reservedSeatsMap)
	//Possible list of families that can be seated
	families := (N - len(reservedSeatsMap)) * 2

	for _, reservedSeats := range reservedSeatsMap {
		flag := false

		if !contains(reservedSeats, "B") && !contains(reservedSeats, "C") && !contains(reservedSeats, "D") && !contains(reservedSeats, "E") {
			flag = true
			families += 1
		}

		if !contains(reservedSeats, "F") && !contains(reservedSeats, "G") && !contains(reservedSeats, "H") && !contains(reservedSeats, "J") {
			flag = true
			families += 1
		}
		if !flag && !contains(reservedSeats, "D") && !contains(reservedSeats, "E") && !contains(reservedSeats, "F") && !contains(reservedSeats, "G") {
			families += 1
		}
	}
	return families
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(Solution(2, "1A 2F 1C")) // 2
	fmt.Println(Solution(1, ""))         //2
	fmt.Println(Solution(40, "1A 3C 2B 40G 5A"))
}
