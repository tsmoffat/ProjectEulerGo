package main

import "fmt"

func main() {
	n := 600851475143
	i := 2
	largeFact := 0
	for i*i <= n {
		if n%i == 0 {
			n = n / i
			largeFact = i
		} else {
			i++

		}
	}
	if n > largeFact {
		largeFact = n
	}
	fmt.Println(largeFact)
}
