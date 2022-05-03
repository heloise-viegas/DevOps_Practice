package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	//or	const conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Println("Welcome to the", conferenceName, "booking application.")
	//fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are available for booking.")
	fmt.Printf("We have a total of %v tickets and %v are available for booking.\n", conferenceTickets, remainingTickets)
	fmt.Println("You can get you're tickets here!")

	//arrays need to have a size defined i.e static size therefore not efficient hence we use Slice
	//var bookings =[50] string {} or var bookings [50]string

	//slice
	bookings := []string{}
	//or var bookings=[]string{} or  var bookings []string for empty slice
	for {
		var fName string
		var lName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name :")
		fmt.Scan(&fName)
		fmt.Println("Enter your last name :")
		fmt.Scan(&lName)
		fmt.Println("Enter your email Id :")
		fmt.Scan(&email)
		fmt.Println("How many tickets do you need ? ")
		fmt.Scan(&userTickets)

		isValid := userTickets <= remainingTickets
		//check for invalid user name etc
		if isValid {

			remainingTickets = remainingTickets - userTickets

			//bookings[0] = fName + " " + lName //for array index has to be known
			bookings = append(bookings, fName+" "+lName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confermation mail at %v\n", fName, lName, userTickets, email)
			fmt.Printf("Remaining tickets are : %v\n", remainingTickets)

			firstNames := []string{}
			//for index,booking:=range bookings{  //index is needed by range but its alos nver used hence use blank identifier _
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				//	var firstName=names[0]
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("All bookings :%v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("All bookings are full!")
				break
			}

		} else {

			//check which condition fails and display text

			fmt.Printf("We have only %v tickets left ,hence we can't book %v tickets.\n", remainingTickets, userTickets)
			//continue //not neede now cause default next line is the loop to be executed
		}

	}
}
