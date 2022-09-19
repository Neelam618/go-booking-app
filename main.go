package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Println("Welcome to", conferenceName, "booking app")
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Printf("conferenceTickets is %T, conferenceTickets is %T, conferenceName is %T.\n", conferenceTickets, conferenceTickets, conferenceName)

	for { ///infinite for loop untill break
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName) //& points to memory address of firstName           // Scan asks user for firstName
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		//user input validation
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")  //if string contains @ character
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidEmail && isValidName && isValidTicketNumber {
			// book ticket in system
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confiramation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// print only first names
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking) //splits slice     names = ["firstname", "lastname"]
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of the bookings are %v\n", firstNames)

			//if all tickets are booked
			if remainingTickets == 0 {
				//end loop
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			fmt.Println("Your input data is invalid, try again")

			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
		}
	}

}
