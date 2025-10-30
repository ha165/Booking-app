package main

import "fmt"

func main() {
 	var confrenceName string  = "Go Confrence"
 	const confrenceTickets int  = 50
 	var remainingTickets uint = 50

	var bookings [50]string
	bookings[0] = "Sajid"
	bookings[1] = "Ahmed"
	

	fmt.Printf("Welcome to %v booking Application!\n", confrenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confrenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend,")

	var userName string
	var userTickets uint
	var userEmail string


	fmt.Println("Enter your name:")
	fmt.Scan(&userName)

	fmt.Println("Enter your email:")
	fmt.Scan(&userEmail)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email shortly.\n", userName, userTickets)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, confrenceName)
}