package main

import "fmt"

func reverse_int(n int) int {
	new_int := 0
	for n > 0 {
		remainder := n % 10
		new_int *= 10
		new_int += remainder
		n /= 10
	}
	return new_int
}
func main() {
	result := 0
	bigPal := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			result = i * j
			if result == reverse_int(result) && result > bigPal {
				bigPal = result
			}
		}
	}
	fmt.Println(bigPal)
}
