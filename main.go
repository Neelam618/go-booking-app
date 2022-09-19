package main

import (
	"fmt"
	"strings"
)

func greetUsers(confName string, confTickets int, remTickets uint) {
	// fmt.Println("Welcome to", confName, "booking app")
	fmt.Printf("Welcome to %v booking app\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confTickets, remTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking) //splits slice     names = ["firstname", "lastname"]
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") //if string contains @ character
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber

}

func getUserInput() (string, string, string, uint) {
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
	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint, conferenceName string, remainingTickets uint, bookings []string) []string {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confiramation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	return bookings
}

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for { ///infinite for loop untill break
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {
			bookTicket(firstName, lastName, email, userTickets, conferenceName, remainingTickets, bookings)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of the bookings are %v\n", firstNames)

			//if all tickets are booked
			if remainingTickets == 0 {
				//end loop
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered is not valid")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			fmt.Println("Your input data is invalid, try again")

			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
		}
	}
}
