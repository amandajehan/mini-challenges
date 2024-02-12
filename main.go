package main

import (
	"fmt"
	"math"
)

func main() {
	n := 64 // dynamic input

	for i := 1; i <= n; i++ {
		num := float64(i)
		numSqrt := math.Sqrt(num)
		numCbrt := math.Cbrt(num)
		var isSquare bool = false
		var isCube bool = false

		if numSqrt == float64(int(numSqrt)) {
			isSquare = true
		}

		if numCbrt == float64(int(numCbrt)) {
			isCube = true
		}

		switch {
		case isSquare && isCube:
			fmt.Println("SquareCube")
		case isSquare && !isCube:
			fmt.Println("Square")
		case !isSquare && isCube:
			fmt.Println("Cube")
		default:
			fmt.Println(num)
		}
	}
}
