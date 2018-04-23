package main

import (
	"fmt"
	"math"
)

func main() {
	n := 0
	pow2 := int(math.Pow(2, 4))
	pow3 := int(math.Pow(3, 2))
	n = pow2 * pow3 * 5 * 7 * 11 * 13 * 17 * 19
	fmt.Println(n)
}
