#https://tour.golang.org/flowcontrol/8
#Exercise: Loops and Functions
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z, prev float64 = 1.0, 1.0
	const DELTA float64 = 1e-9
	for {
		z = prev - ((math.Pow(prev,2) - x) / (2 * prev))
		if math.Abs(prev - z) < DELTA {
			return z
		}
		prev = z
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
