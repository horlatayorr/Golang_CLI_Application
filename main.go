package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // dynamic lists using slices

type UserData struct { //struct allows to create custom data types and storing different data types without conversion unlike maps
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTickets {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)                                              //add 1 to wait group
		go sendTicket(userTickets, firstName, lastName, email) //concurrency using go keyword

		firstNames := getFirstNames()
		fmt.Printf("The First names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("All tickets are sold out. Thank you for booking")
		}
		//break
	} else {
		if !isValidName {
			fmt.Printf("First Name or Last Name you entered is too short\n")
		}
		if !isValidEmail {
			fmt.Printf("Email you entered doesn't contain '@' sign or '.'\n")
		}
		if !isValidTickets {
			fmt.Printf("Number of Tickets you entered is invalid\n")
		}
	}
	wg.Wait() //wait for all goroutines to finish
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames() []string { //input parameters inside parenthesis, output parameter type outside parenthesis
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
	var userTickets uint

	// ask user for their name
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Printf("Enter your number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) uint {
	remainingTickets := remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v remaining tickets for %v\n", remainingTickets, conferenceName)

	return remainingTickets

	//var userData = make(map[string]string) //definition of map is for key value pairs
	//userData["firstName"] = firstName
	//userData["lastName"] = lastName
	//userData["email"] = email
	//userData["numberOfTickets"] = fmt.Sprint(userTickets) //convert uint to string does the same as strconv.FormatUint(uint64(userTickets),10)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("###############")

	wg.Done() //decrement 1 from wait group
}
