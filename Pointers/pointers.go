package main

import "fmt"

func main() {
	age := 32

	agePointer := &age

	fmt.Println("Age: ", *agePointer)

	getAdultYears(agePointer)

	fmt.Println("Adult year:", *agePointer)

	fmt.Println("Age: ",  age)

}

func getAdultYears(age *int)  {
	*age = *age - 18
}
