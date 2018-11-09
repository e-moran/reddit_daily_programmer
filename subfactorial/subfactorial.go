package main

import (
	"fmt"
	"math"
)

func main() {
	var n, subfact int
	fmt.Println("Enter an Integer:")
	fmt.Scanf("%d", &n)
	subfact = subfactorial(n)
	fmt.Println(subfact)
}

//Recursively compute subfactorial
func subfactorial(n int) int {
	if n == 1 {
		return 0
	}

	return n*subfactorial(n-1) + int(math.Pow(-1, float64(n)))
}