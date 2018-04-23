package main

import "fmt"

func main() {
	sum := 0
	curr := 1
	prev := 0
	temp := 0
	for curr < 4000000 {
		temp = curr
		curr = curr + prev
		prev = temp
		if curr%2 == 0 {
			sum += curr
		}
	}
	fmt.Println(sum)
}
