package main

import (
	"errors"
	"fmt"
	"os"
)

const textFileName = "userInput.txt"

func main() {
	readuserInput()
	revenue, revError := getUserInput("Revenue: ")
	printUserError(revError)

	expenses, exError := getUserInput("\nExpenses: ")
	printUserError(exError)

	taxRate, taxError := getUserInput("\nTax Rate: ")
	printUserError(taxError)

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	saveUserInput(ebt, profit, ratio)
	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	var returnError error
	if userInput < 0 {

		userInput = 0

		returnError = errors.New("failed to capture user input with negative value")
	}
	return userInput, returnError
}
func printUserError(inputError error) {
	if inputError != nil {
		fmt.Print(inputError.Error())
	}
}
func saveUserInput(ebt, profit, tax float64) {
	saveText := fmt.Sprint(ebt, profit, tax)

	os.WriteFile(textFileName, []byte(saveText), 0644)
}
func readuserInput() {
	var data, error = os.ReadFile(textFileName)
	if(error != nil){
		fmt.Println(error.Error())
		return
	}

	var outputText = string(data)

	fmt.Println(outputText)
}
