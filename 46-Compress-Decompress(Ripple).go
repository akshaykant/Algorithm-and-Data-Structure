/*

Objective: Write an algorithm to encode and decode a given string by
using the count of repeated characters (you can think of this as a simple
compression algorithm).
Example:
encode("AAABB") = "A3B2"
decode("A3B2") = "AAABB"

You need to implement two functions for a simple compression algorithm:

Encode: Compress a given string by using the count of repeated characters.
Decode: Decompress the string back to its original form.
Example:
Encoding: "AAABB" should be encoded as "A3B2".
Decoding: "A3B2" should be decoded back to "AAABB".

*/

/*
Edge Case - to Discuss

encoding needs to have edge case -where compression input should not increase the size.
Ex. AB should not give A1B1
*/

package main

import (
    "fmt"
    "strconv"
)

// encode compresses the string unless it would make the string longer.
func encode(input string) string {
    if len(input) == 0 {
        return ""
    }

    result := ""
    for i := 0; i < len(input); i++ {
        count := 1
        // Count occurrences of the same character
        for i+1 < len(input) && input[i] == input[i+1] {
            i++
            count++
        }

        // Append the character and count, only if count > 1
        if count > 1 {
            result += string(input[i]) + strconv.Itoa(count)
        } else {
            result += string(input[i])
        }
    }
    return result
}

// decode expands the encoded string back to its original format.
func decode(input string) string {
    result := ""
    var numStr string

    for i := 0; i < len(input); i++ {
        char := input[i]
        if char >= '0' && char <= '9' {
            numStr += string(char)
        } else {
            if numStr != "" {
                count, _ := strconv.Atoi(numStr)
                result += repeatLastChar(result, count)
                numStr = ""
            }
            result += string(char)
        }
    }

    // Handle the last number if it exists
    if numStr != "" {
        count, _ := strconv.Atoi(numStr)
        result += repeatLastChar(result, count)
    }

    return result
}

// repeatLastChar repeats the last character of a string count times.
func repeatLastChar(s string, count int) string {
    if len(s) == 0 {
        return ""
    }
    lastChar := s[len(s)-1]
    result := ""
    for i := 0; i < count; i++ {
        result += string(lastChar)
    }
    return result
}

func main() {
    original := "AB"
    encoded := encode(original)
    decoded := decode(encoded)
    fmt.Printf("Original: %s, Encoded: %s, Decoded: %s\n", original, encoded, decoded)
}
