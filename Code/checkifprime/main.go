package main

import (
	"fmt"
	"math"
)

func main()  {
	var inNum int
	fmt.Println("Enter a whole number:")
	fmt.Scanln(&inNum)
	isPrime := checkPrim(inNum)
	fmt.Println("Is the number Prime: ", isPrime)
}

func checkPrim(innm int) bool {
	if innm == 1 || innm == 0 {
		fmt.Println("The number is not prime")
		return false
	}
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(innm)))); i++ {
		if innm % i == 0 {
			fmt.Printf("This number %v is not prime: \n", innm)
			return false
		}
	}
	return true
}