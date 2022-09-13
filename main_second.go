// same package different file
package main

import (
	"fmt"
	"sync"
	"time"

	"mypath.com/go_test_module/helper"
)

// global var
// wait group (wait for thread to finish)
var wg = sync.WaitGroup{}

func testsecond() {
	fmt.Printf("june helloworld\n")
}

func Sendemail(firstname string, email string) {
	time.Sleep(20 * time.Second) // wait 20 sec
	var ticket = fmt.Sprintf("user ticket for %v", firstname)
	fmt.Println("########")
	fmt.Printf("Send ticket for %v, to email address %v\n", ticket, email)
	fmt.Println("########")
	wg.Done()
}

// defining struct
type UserData struct {
	firstname string
	lastname  string
	email     string
	city      string
}

func Kiosk_function_as_struct(current_ticket_no int) int {
	var database = make([]UserData, 1) // List of maps - increases size automatically

	for current_ticket_no > 0 {

		firstname, lastname, email, city, ticket_no := get_userinput()
		var userdata = UserData{
			firstname: firstname,
			lastname:  lastname,
			email:     email,
			city:      city,
		}
		fmt.Printf("%v \n ", userdata)
		database = append(database, userdata)

		if current_ticket_no > 50 {
			fmt.Printf("cannot buy more than 50\n")
			break
		}
		if user_input_validation(firstname, lastname, email, ticket_no) == true {
			fmt.Printf("valid input..\n")
		} else {
			fmt.Printf("not a valid input..\n")
			break
		}
		current_ticket_no = current_ticket_no - ticket_no
		fmt.Printf("-----summary----- \nfirstname: %v  lastname: %v  email: %v  remaining tickets: %v\n", firstname, lastname, email, current_ticket_no)

		// concurrency in GO
		wg.Add(1)
		go Sendemail(firstname, email)

		// city check
		if helper.Check_valid_city(city) == false {
			fmt.Printf("not a valid city exiting for loop")
			break
		}
		fmt.Printf("\n")

	}
	wg.Wait()
	fmt.Printf("...", database)
	return current_ticket_no
}
