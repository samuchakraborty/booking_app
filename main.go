package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go Conference"
	const conferenceTicket int = 50
	var remainingTicket int = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v ticket and %v are still available\n", conferenceTicket, remainingTicket)

	fmt.Printf("Get your Ticket here\n")

	// creating array
	// var booking [50]string

	///
	// var booking []string

	// var booking []string
	// booking = append(booking, "SAMU chakraborty")
	// booking = append(booking, "RAM chakraborty")

	// firstNames := []string{}

	//fmt.Println(firstName)
	// fmt.Println(booking)
	// fmt.Println(booking2[1])
	// fmt.Println(len(booking2))

	// fmt.Printf("type is %T\n", booking2)

	for remainingTicket > 0 || remainingTicket < 50 {

		var firstName string
		var lastName string

		var email string
		var userTicket int

		var booking2 []string
		fmt.Printf("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Printf("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Printf("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Printf("Enter your ticket: ")

		fmt.Scan(&userTicket)

		isValidateName := len(firstName) > 2 && len(lastName) > 2

		isValidateEmail := strings.Contains(email, "@")

		isValidateTicketNumber := userTicket > 0 && userTicket <= remainingTicket

		if isValidateName && isValidateEmail && isValidateTicketNumber {

			remainingTicket = remainingTicket - userTicket
			booking2 = append(booking2, firstName+" "+lastName)

			fmt.Printf("user %v %v is get %v ticket. \n", firstName, lastName, userTicket)
			fmt.Printf("thank you and an invitation is send to your email address at %v.\n", email)
			firstNames := []string{}

			// for i := 0; i < len(booking2); i++ {
			// 	//fmt.Println(strings.Split(booking2[i], " "))
			// 	var name = strings.Split(booking2[i], " ")
			// 	//fmt.Println("name : ", name)
			// 	firstNames = append(firstNames, name[0])

			// }
			for _, booking := range booking2 {
				var name = strings.Fields(booking)
				firstNames = append(firstNames, name[0])
			}

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

		
		//  var booking = [50]string{"Nan ", "Josh", "POsh"}

	}

}
