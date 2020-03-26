package main

import (
	"fmt"
)

func main() {
	var amount float32
	keep := true
	var continueInput string
	for keep != false {
		fmt.Println("Type a price in euros and convert it in Dollars, Yens and Pounds:")
		fmt.Scan(&amount)
		dollars := amount * 1.1
		yen := amount * 120.77
		pounds := amount * 0.91
		fmt.Println(fmt.Sprintf("Euros: %.2f \n", amount))
		fmt.Println(fmt.Sprintf("Dollars: %.2f \n", dollars))
		fmt.Println(fmt.Sprintf("Yens: %.2f \n", yen))
		fmt.Println(fmt.Sprintf("Pounds: %.2f \n", pounds))
		fmt.Println("Do you want to convert another price ? (Y/n)")
		fmt.Scan(&continueInput)
		if continueInput != "Y" {
			keep = false
		}
	}
}