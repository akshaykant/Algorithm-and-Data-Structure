//Shuffle a set of numbers without duplicates.
package main

import (
	"fmt"
	"math/rand"
	"time"

)

func main(){
	in := []int{1,2,3,4,5,6,7,8,9,10}

	out := shuffle(in)

	fmt.Println(out)
}

//Knuth Shuffle
//In iteration i, pick integer r between o and i at random.
//Fisher Yates proposed that Knuth shuffle produces a uniformly random permutation of the input array in linear time.
func shuffle(in []int) []int{
	
	//Use better seed for better randomness.
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(in); i = i + 1{
		r := rand.Intn(i + 1)
		in[i], in[r] = in[r], in[i]
	}

	return in
}