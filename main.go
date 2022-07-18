package main

import (
	"BOOKING_APP/helper"
	"fmt"
	// "strconv"
	// "strings"
)

//package level variable
var conferenceName = "Go Conference"

const conferenceTicket int = 50

var remainingTicket int = 50
// var booking = make([]map[string]string, 0)
var booking = make([]UserData, 0)


type UserData struct {
	firstName  string
	lastName   string
	email      string
	userTicket int
}

func main() {

	greetUsers()

	//for remainingTicket > 0 || remainingTicket < 50

	for {

		firstName, lastName, email, userTicket := getUserInput()
		isValidateName, isValidateEmail, isValidateTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTicket)

		if isValidateName && isValidateEmail && isValidateTicketNumber {

			bookTicket(userTicket, firstName, lastName, email)
			firstNames := getFirstName()

			fmt.Printf("The booking user is %v\n", firstNames)

			if remainingTicket == 0 {
				fmt.Printf("Our conference tickets are sold. Thank you and try in next year")
				break

			} else {
				fmt.Printf("Remaining ticket are %v \n", remainingTicket)

			}

		} else {
			//fmt.Printf("we only have %v remaining ticket, but you can't book %v tickets\n", remainingTicket, userTicket)
			if !isValidateName {
				fmt.Println("The first name or last name you entered is too short")
			}
			if !isValidateEmail {
				fmt.Println("The email address you entered doesn't contain @ sign")

			}
			if !isValidateTicketNumber {
				fmt.Println("number of ticket you entered is invalid")
			}

			//fmt.Println("Your input are wrong, please try again later")

			continue
		}

	}

}

func greetUsers() {

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v ticket and %v are still available\n", conferenceTicket, remainingTicket)

	fmt.Printf("Get your Ticket here\n")

}

func getFirstName() []string {
	firstNames := []string{}

	// for i := 0; i < len(booking); i++ {
	// 	//fmt.Println(strings.Split(booking[i], " "))
	// 	var name = strings.Split(booking[i], " ")
	// 	//fmt.Println("name : ", name)
	// 	firstNames = append(firstNames, name[0])

	// }
	for _, bookings := range booking {

		firstNames = append(firstNames, bookings.firstName)
		// fmt.Printf("The booking user is %v\n", firstNames)
	}
	return firstNames

}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string

	var email string
	var userTicket int

	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Printf("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Printf("Enter your ticket: ")

	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicket(userTicket int, firstName string, lastName string, email string) {

	remainingTicket = remainingTicket - userTicket
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTicket"] = strconv.Itoa(userTicket)

	var userData = UserData{

		firstName: firstName,
		lastName : lastName,
		email: email,
		userTicket: userTicket,

	}


	booking = append(booking, userData)
	fmt.Printf("UserData is: %v\n", userData)

	fmt.Printf("List of user is: %v\n", booking)

	fmt.Printf("user %v %v is get %v ticket. \n", firstName, lastName, userTicket)
	fmt.Printf("thank you and an invitation is send to your email address at %v.\n", email)
	//fmt.Printf("Remaining ticket are %v \n", remainingTicket)

}

// func switchCaseGo()  {

// 	city := "Dhaka"

// 	switch city {

// 	case "London":
// 		fmt.Println("City is London")
// 	case "Bangladesh", "Dhaka":
// 		fmt.Println("City in Bangladesh")
// 	default:
// 		fmt.Println("City is not found")
// 	}

// }
