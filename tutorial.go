package main

import "fmt"

func main() {
	fmt.Println("Welcome to the quiz game!")
	var name string
	fmt.Println("Hello sweet chile, what is ya name?")
	fmt.Scan(&name)
	fmt.Printf("Nice to meet you %v! welcome to the THUNDERDOME BROTHER!\n", name)
	fmt.Printf("How many years have you seen in your lifetime?")
	var age int
	fmt.Scan(&age)
	if age >= 10 {
		fmt.Println("You have been accepted")
	}
	if age < 10 {
		fmt.Println("You cannot participate")
	}

}
