package main

import "fmt"

func main() {
 	var confrenceName string  = "Go Confrence"
 	const confrenceTickets int  = 50
 	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking Application!\n", confrenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confrenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend,")

	var userName string
	var userTickets uint

	fmt.Println("Enter your name:")
	fmt.Scan(&userName)

	userTickets = 2

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email shortly.\n", userName, userTickets)
}