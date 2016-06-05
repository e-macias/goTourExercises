//https://tour.golang.org/flowcontrol/8
//Exercise: Loops and Functions
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


//https://tour.golang.org/moretypes/18
//Exercise: Slices
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for i := range a {
		a[i] = make([]uint8, dx)
		for j := range a[i] {
			a[i][j] = uint8((i^j)/2+(i^j))
		}
	}
	return a
}

func main() {
	pic.Show(Pic)
}


//https://tour.golang.org/moretypes/23
//Exercise: Maps
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordsMap := make(map[string]int)
	for _, key := range strings.Fields(s) {
		count, ok := wordsMap[key]
		if ok == true {
			count++
			wordsMap[key] = count //++count not allowed
		} else {
			wordsMap[key] = 1
		}
	}
	return wordsMap
}

func main() {
	wc.Test(WordCount)
}
