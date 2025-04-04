package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"

func readBalance() (float64, error) {
	data, err := os.ReadFile(balanceFile)
	if err != nil {
		return 0, errors.New("can not read balance")
	}
	dataString := string(data)
	balance, err := strconv.ParseFloat(dataString, 64)
	if err != nil {
		return 0, errors.New("can not convert to a float")
	}
	return balance, nil
}

func writeBalance(balance float64) {
	balanceToText := fmt.Sprint(balance)
	os.WriteFile(balanceFile, []byte(balanceToText), 0644)
}

func main() {
	var accountBalance, err = readBalance()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What would you like to do today?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("You balance is: R%.2f\n", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Deposit amount: R")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invali amount, deposit amount must be greater then R0")
				continue
			}
			accountBalance += depositAmount
			fmt.Printf("Updated balence: R%.2f\n", accountBalance)
			writeBalance(accountBalance)
		case 3:
			var withdralAmount float64
			fmt.Print("Withdraw amount: R")
			fmt.Scan(&withdralAmount)
			if withdralAmount <= 0 {
				fmt.Println("Invali amount, withdrawl amount must be greater then R0")
				continue
			}
			if withdralAmount > accountBalance {
				fmt.Println("Invali amount, withdrawl greater then balence")
				continue
			}
			accountBalance -= withdralAmount
			fmt.Printf("Updated balence: R%.2f\n", accountBalance)
			writeBalance(accountBalance)
		default:
			fmt.Println("Exiting app.")
			fmt.Println("Thank you for using Go Bank!")
			return
		}
	}
}
