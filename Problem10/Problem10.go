package main

import (
	"fmt"
	"math"
)

func SieveOfEratosthenes(value int) {
	sum := 0
	f := make([]bool, value)
	for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
		if f[i] == false {
			for j := i * i; j < value; j += i {
				f[j] = true
			}
		}
	}
	for i := 2; i < value; i++ {
		if f[i] == false {
			sum += i
		}
	}
	fmt.Println("%v", sum)
}
func main() {
	SieveOfEratosthenes(2000000)
}
