package main

import (
	"fmt"
	"strings"
)

func main() {
	var confrenceName string = "Go Confrence"
	const confrenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking Application!\n", confrenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confrenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend,")

	for {
		var firstName string
		var lastName string
		var userTickets uint
		var userEmail string

		fmt.Println("Enter your First Name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your Last Name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&userEmail)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email shortly.\n", firstName, userTickets)
		fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, confrenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}
		fmt.Printf("These are all our bookings: %v\n", bookings)
	}
}
