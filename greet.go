package main

import (
	"fmt"
)

func greet(customerName string) {
	fmt.Printf("%52s %s\n", "Namaste ", customerName)
	fmt.Printf("%72s\n", "_/\\_ Welcome to Sharma Bhojanalya! _/\\_ ")
	fmt.Println()
}

func sayTata(customerName string) {
	fmt.Println()
	fmt.Printf("%17s", " ")
	fmt.Printf("_/\\_ Thank you %v for visiting Sharma Bhojanalya! _/\\_\n", customerName)
	fmt.Printf("%55s\n", "We hope to see you again!")
	fmt.Printf("%51s\n", "Aapka din shubh rahe !! ")
	fmt.Printf("%59s\n", " ----------------------------------- ")
	fmt.Printf("%59s\n", "|      Credit :- Rohit Parmar        |")
	fmt.Printf("%59s\n", " ----------------------------------- ")
}
