package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}
	fmt.Println("%d", sum)
}
