package main

import(
	"fmt"
)

func main() {
	c := "#"
	fmt.Printf(c+"\n")
	for i := 0; i <= 10; i++ {
		c = c + "#"
		fmt.Printf(c+"\n")
	}

}