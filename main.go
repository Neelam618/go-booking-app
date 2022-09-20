package main

import (
	"fmt"
	"go-booking-app/helper"  
)

const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)       //[{}]

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", remainingTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { //splits slice     names = ["firstname", "lastname"]
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets
	//create map for user
	var userData = UserData	{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings )
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confiramation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func main() {
	greetUsers()
	for { ///infinite for loop untill break
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {
			bookTicket(firstName, lastName, email, userTickets)

			firstNames := getFirstNames()
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
		}
	}
}
