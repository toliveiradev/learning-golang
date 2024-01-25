package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName            string
	lastName             string
	email                string
	numberOfTickets      uint
	isOptedForNewsletter bool
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	firstName, lastName, email, userTicketsCount := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTicketsCount)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTicketsCount, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTicketsCount, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Printf("Sorry, %v is sold out.\n", conferenceName)
		}
	} else {
		if !isValidTicketNumber {
			fmt.Printf("Sorry, only %v tickets are remaining. So you can't book %v tickets\n", remainingTickets, userTicketsCount)
		}
		if !isValidName {
			fmt.Println("First name or last name you entered is too short.")
		}
		if !isValidEmail {
			fmt.Println("Email you entered doesn't contain @ sign.")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and only %v tickets are remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicketsCount uint

	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want to book: ")
	fmt.Scan(&userTicketsCount)

	return firstName, lastName, email, userTicketsCount
}

func bookTicket(userTicketsCount uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTicketsCount

	var userData = UserData{
		firstName:            firstName,
		lastName:             lastName,
		email:                email,
		numberOfTickets:      userTicketsCount,
		isOptedForNewsletter: true,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)
	/*
		fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Slice type: %T\n", bookings)
		fmt.Printf("Slice lenght: %v\n", len(bookings))
	*/
	fmt.Printf("Thank you %v %v for booking %v tickets for %v. You will get a confirmation at %v\n", firstName, lastName, userTicketsCount, conferenceName, email)
	fmt.Printf("There are %v tickets remaining for %v. \n", remainingTickets, conferenceName)
}

func sendTicket(userTicketsCount uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicketsCount, firstName, lastName)
	fmt.Println("########################")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("########################")
	wg.Done()
}
