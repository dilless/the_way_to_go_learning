package main

import "fmt"

func main() {
	var i int = 5

	for {
		i -= 1
		fmt.Printf("The variable i is now: %d\n", i)
		if i < 0 {
			break
		}
	}
}
