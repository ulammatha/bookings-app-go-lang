package main

import (
	"fmt"
	"strings"
)

/*
NOTES:
	import
	package
	:= is same as var conferenceName string = "Go Conference", this is only for "var" not for "const" also we cannot define "explicit type" when we use :=
	fmt.Printf
	fmt.Println
	fmt.Scan
	strings.Fields(booking) // slice the string based on spaces
	strings.Contains(email, "@")
	arrays, slices
		var bookings [50]string // ARRAY usage
		bookings := []string{} // SLICES usage with empty object
	append
	for
	if
	switch city {
		case "New York":
			fmt.Println("Booking for New York conference")
		case "London", "Berlin":
			fmt.Println("Booking for London or Berlin conference")
		default:
			fmt.Println("No valid city selected")
	}
	func
	package level variables -> variables that can be accessed between functions, package level variables cannot use :=, need to use var
*/

var conferenceName = "Go Conference" // package level  variable
const conferenceTickets uint = 50    // package level  variable
var remainingTickets uint = 50       // package level  variable
var bookings = []string{}            // SLICES usage, empty. // package level  variable

// entry point of the application
func main() {
	greetUsers()
	for remainingTickets > 0 && uint(len(bookings)) < conferenceTickets {
		firstName, lastName, email, userTickets := getUserInput() // local variables
		isValidName, isValidEmail, isValidTicketNumber := validateUserInputs(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("The first name of bookings are: %v\n", firstNames)

			// var hasTickets bool = remainingTickets == 0
			hasTickets := remainingTickets == 0
			if hasTickets {
				// end the program
				fmt.Printf("Our conference is sold out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("First name or last name you entered is too short (minimum 2 characters required)\n")
			}
			if !isValidEmail {
				fmt.Printf("Email entered is invalid, try again.\n")
			}
			if !isValidTicketNumber {
				fmt.Printf("The number of tickets you requested is invalid, try again.\n")
			}
		}
	}
}
func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("conferenceName is %T \nconferenceTickets is %T \nremainingTickets is %T \n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { // _ ignore index since no use here, in GO _ means ignore it
		var names = strings.Fields(booking) // slice the string based on spaces
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInputs(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Enter First Name:")
	fmt.Scan(&firstName) // pointer variable to store user typed input

	fmt.Printf("Enter Last Name:")
	fmt.Scan(&lastName) // pointer variable to store user typed input

	fmt.Printf("Enter your Email:")
	fmt.Scan(&email) // pointer variable to store user typed input

	fmt.Printf("Enter number of Tickets to Book:")
	fmt.Scan(&userTickets) // pointer variable to store user typed input
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName // ARRAY usage
	bookings = append(bookings, firstName+" "+lastName) // SLICES usage
	fmt.Printf("Thank you %v  %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remainingTickets for %v \n", remainingTickets, conferenceName)
	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The First value: %v\n", bookings[0])
	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("Slice length: %v\n", len(bookings))
	// return bookings
}

// in terminal,  "go run main.go"
