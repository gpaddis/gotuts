package main

import (
	"fmt"
)

func main() {
	numbers := make([]int, 11)
	for n := 0; n < len(numbers); n++ {
		if n%2 == 0 {
			fmt.Printf("%d is even\n", n)
		} else {
			fmt.Printf("%d is odd\n", n)
		}
	}
}
