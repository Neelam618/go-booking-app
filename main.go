package main

import "fmt"

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var userTickets uint

	// fmt.Println("Welcome to", conferenceName, "booking app") OR
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Printf("conferenceTickets is %T, conferenceTickets is %T, conferenceName is %T.\n", conferenceTickets, conferenceTickets, conferenceName )
	
	var firstName string
	var lastName string
	var email string

	userTickets = 2
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)  //& points to memory address of firstName           // Scan asks user for firstName
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName) 
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)  

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confiramation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
