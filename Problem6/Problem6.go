package main

import (
	"fmt"
	"math"
)

func main() {
	sumsq := 0
	sqsum := 0
	for i := 1; i <= 100; i++ {
		sumsq += int(math.Pow(float64(i), 2))
		sqsum += i
	}
	sqsum = int(math.Pow(float64(sqsum), 2))
	diff := sqsum - sumsq
	fmt.Println(diff, sqsum, sumsq)
}
