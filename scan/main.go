package main

import (
	"fmt"
)

func main() {
    var input string
    
    for input != "exit" {
        fmt.Print("Enter Something or type 'exit' to close the program: \n")
        fmt.Scanln(&input)
        fmt.Print("You typed: "+input+"\n")
    }

}