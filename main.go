package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < 2; i++ {
		var accountBalence = 3000.00
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What would you like to do today?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Printf("You balance is: R%.2f\n", accountBalence)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Print("Deposit amount: R")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invali amount, deposit amount must be greater then R0")
				return
			}
			accountBalence += depositAmount
			fmt.Printf("Updated balence: R%.2f\n", accountBalence)
		} else if choice == 3 {
			var withdralAmount float64
			fmt.Print("Withdraw amount: R")
			fmt.Scan(&withdralAmount)
			if withdralAmount <= 0 {
				fmt.Println("Invali amount, withdrawl amount must be greater then R0")
				return
			}
			if withdralAmount > accountBalence {
				fmt.Println("Invali amount, withdrawl greater then balence")
				return
			}
			accountBalence -= withdralAmount
			fmt.Printf("Updated balence: R%.2f\n", accountBalence)
		} else if choice == 4 {
			fmt.Println("Exiting app.")
			os.Exit(0)
		}

	}
}
