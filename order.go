package main

import (
	"fmt"
	"strings"
)

// order using this function
func orderItems() {
	printMenu()
	var itemNumber int
	var noOfPlates int

	for {
		fmt.Println()
		fmt.Println("Enter '0' to exit.")
		fmt.Print("Enter the serial no. of the item to order: ")

		fmt.Scan(&itemNumber)
		if itemNumber == 0 {
			break
		}
		if itemNumber > 13 || itemNumber < 0 {
			fmt.Print(" -------------------------------\n")
			fmt.Print("|        Invalid Choice         |\n")
			fmt.Print("|     Please Choose again !!    |\n")
			fmt.Print(" -------------------------------\n")
			continue
		}
		var choiceName string
		var itemPrice float64

		for index, item := range menu {
			if index+1 == int(itemNumber) {
				choiceName = item.itemName
				itemPrice = item.itemPrice
				break
			}
		}
		fmt.Printf("How many %v do you want: ", choiceName)
		fmt.Scan(&noOfPlates)
		if noOfPlates == 0 {
			continue
		} else {
			for index /*, item*/ := range menu {
				if index+1 == int(itemNumber) { //convert optionNumber into int becoz it is of type int

					customerOrder[choiceName] += noOfPlates
					if customerOrder[choiceName] <= 0 {
						delFromOrder(itemNumber)
					}
					//alternative way
					//customerOrder[item.itemName] += noOfPlates // adding customerOrder[item.itemName] again in case you order that item again
					subTotalBill += itemPrice * float64(noOfPlates)
					//alternative way
					//subTotalBill += item.itemPrice*noOfPlates
					break
				}
			}
			if noOfPlates < 0 {
				fmt.Printf("\nYou just Cancelled %v %v which amounts to ₹%v.\n", noOfPlates, choiceName, itemPrice*float64(noOfPlates))
			} else {
				fmt.Printf("\nYou just ordered %v %v which amounts to ₹%v.\n", noOfPlates, choiceName, itemPrice*float64(noOfPlates))
			}
			//print what you ordered till now
			orderTillNow()
		}
		fmt.Println()
	}
}

// print it everytime you add an item
func orderTillNow() {
	//Print what you've ordered till now
	fmt.Println("\nYour order till now: ")
	fmt.Printf("%s\n", strings.Repeat("-", 32))
	fmt.Printf(" %-12s %s\n", "Quantity", "Item")
	fmt.Printf("%s\n", strings.Repeat("-", 32))

	for i := range customerOrder {
		fmt.Printf(" %-12v %v\n", customerOrder[i], i)
	}

	fmt.Printf("%s\n", strings.Repeat("-", 32))
}
