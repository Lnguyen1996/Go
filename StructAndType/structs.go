package main

import (
	"fmt"

	user "example.com/structs/userStruct"
)

func main() {
	userfirstName := getUserData("Please enter your first name: ")
	userlastName := getUserData("Please enter your last name: ")
	userbirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, error := user.New(
		userfirstName,
		userlastName,
		userbirthdate,
	)

	if error != nil {
		panic(error.Error())
	}
	// ... do something awesome with that gathered data!
	appUser.PrintUserDetails()
	appUser.ClearUserName()

	appUser.PrintUserDetails()

	admin:= user.NewAdmin("test@test.com","test123")

	admin.PrintUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
