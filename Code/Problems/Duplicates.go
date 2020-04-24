package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := [3]string{"me", "me", "me"}
	result := hasDup(ints)
	fmt.Printf("This is the result %v", result)
}

func hasDup(ints [3]string) bool  {
	seen := false
	sort.Sort(ints)
	for i range(0, len(ints)-1){
		if  ints[i] == ints[i+1]{
			 return false
		}
	}
	return seen
}