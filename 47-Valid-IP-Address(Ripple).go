/*
Part - 1

Given a string, write a function that returns whether or not the string is a valid IP address [0-255].[0-255].[0-255].[0-255].
*/

/*
Explanation of the Code

String Splitting: The strings.Split() function is used to divide the IP address into segments by the dot.
Length Check: Ensures there are exactly four segments.
Numeric Conversion and Validation:
The segment is checked for leading zeros using string properties.
strconv.Atoi() attempts to convert the string segment into an integer, ensuring the segment is purely numeric.
It also validates that the number falls within the acceptable range for IPv4 addresses (0 to 255).
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// isValidIP checks if the input string is a valid IPv4 address.
func isValidIP(ip string) bool {
	// Split the IP address into its components.
	parts := strings.Split(ip, ".")
	// An IPv4 address must have exactly four parts.
	if len(parts) != 4 {
		return false
	}

	for _, part := range parts {
		// Check if the part is non-empty and does not have leading zeros.
		if part == "" || (len(part) > 1 && part[0] == '0') {
			return false
		}

		// Convert the string to an integer.
		num, err := strconv.Atoi(part)
		if err != nil {
			// If it's not a number, return false.
			return false
		}

		// Check if the integer is outside the range of [0, 255].
		if num < 0 || num > 255 {
			return false
		}
	}

	return true
}

func main() {
	testIPs := []string{"192.168.1.1", "255.255.255.255", "192.168.1.256", "192.168.1", "192.168.01.1"}
	for _, ip := range testIPs {
		fmt.Printf("%s is a valid IP address: %v\n", ip, isValidIP(ip))
	}
}



/*
Part-2

Given a string, write a function that returns whether or not a valid IP address can be constructed from it with no remaining characters. 

eg. 123123123123 = true 
    55555555555 = false
*/

/*
Approach
Check Length: The length of the string must be between 4 and 12 characters. Fewer than 4 characters won't allow four segments, and more than 12 characters would force one segment to exceed 3 digits.
Recursive Backtracking:
Try to build an IPv4 address by recursively choosing segments from the string.
At each step, extract between 1 and 3 characters (inclusive) from the current position, convert to a number, and check if it's a valid segment (0-255).
If the segment is valid, proceed to extract the next segment.
*/

package main

import (
	"fmt"
	"strconv"
)

// isValidIPAddressFromString checks if a valid IP address can be constructed from the input string.
func isValidIPAddressFromString(s string) bool {
    if len(s) < 4 || len(s) > 12 {
        return false
    }
    var segments []string
    return backtrack(s, 0, &segments)
}

// backtrack tries to construct a valid IP address recursively.
func backtrack(s string, start int, segments *[]string) bool {
    // If we've collected 4 segments and we're at the end of the string, it's a valid IP.
    if len(*segments) == 4 {
        if start == len(s) {
            return true
        }
        return false
    }

    // Try to take 1 to 3 digits to form the next segment.
    for l := 1; l <= 3; l++ {
        if start+l > len(s) {
            break
        }
        segment := s[start : start+l]
        if isValidSegment(segment) {
            *segments = append(*segments, segment)
            if backtrack(s, start+l, segments) {
                return true
            }
            *segments = (*segments)[:len(*segments)-1]
        }
    }
    return false
}

// isValidSegment checks if the string is a valid segment of an IP address.
func isValidSegment(segment string) bool {
    // Leading zero is only allowed if the segment is exactly "0".
    if len(segment) > 1 && segment[0] == '0' {
        return false
    }
    num, err := strconv.Atoi(segment)
    if err != nil {
        return false
    }
    return num >= 0 && num <= 255
}

func main() {
    testStrings := []string{"123123123123", "55555555555", "19216811", "172162541"}
    for _, s := range testStrings {
        fmt.Printf("%s can form a valid IP address: %v\n", s, isValidIPAddressFromString(s))
    }
}

