package main

import (
	"fmt"
	"strings"

	"mypath.com/go_test_module/helper"
)

// article "mypath.com/go_test_module/article"

// functions
func test_function(ticket_no int) (int, bool) { // returning 2 values
	fmt.Printf("left tickets are this much tickets %v \n", ticket_no)
	return ticket_no, true
}

func get_userinput() (string, string, string, string, int) {
	fmt.Print("next person in line .....\n")
	var firstname string
	fmt.Printf("firstname: \n")
	fmt.Scan(&firstname) // need to use & reference
	var lastname string
	fmt.Printf("lastname: \n")
	fmt.Scan(&lastname) // need to use & reference
	var email string
	fmt.Printf("email: \n")
	fmt.Scan(&email) // need to use & reference
	var city string
	fmt.Printf("city: \n")
	fmt.Scan(&city) // need to use & reference
	var ticket_no int
	fmt.Printf("ticket_no: \n")
	fmt.Scan(&ticket_no) // need to use & reference

	return firstname, lastname, email, city, ticket_no

}

func kiosk_function(current_ticket_no int) int {
	var database = make([]map[string]string, 1) // List of maps - increases size automatically

	for current_ticket_no > 0 {

		var userdata = make(map[string]string)
		firstname, lastname, email, city, ticket_no := get_userinput()
		userdata["firstname"] = firstname
		userdata["lastname"] = lastname
		userdata["email"] = email
		userdata["city"] = city
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

		// city check
		if helper.Check_valid_city(city) == false {
			fmt.Printf("not a valid city exiting for loop")
			break
		}
		fmt.Printf("\n")

	}
	fmt.Printf("...", database)
	return current_ticket_no
}

func user_input_validation(firstname string, lastname string, email string, ticket_no int) bool {
	// user input validation
	validuserinput := len(firstname) > 0 && len(lastname) > 0 && len(email) > 0 && strings.Contains(email, "@")
	if validuserinput {
		fmt.Printf("Thank you for booking tickets  %v %v, you have bought %v tickets\n", firstname, lastname, ticket_no)
		return true
	} else if !validuserinput {
		fmt.Printf("user input is wrong .. ")
		return false
	} else {
		fmt.Printf("something else is wrong .. ")
		return false
	}
}

// package level variables
// user input - for loop example
const totalnumber int = 50          // const cannot be changed
var remaining_tickets = totalnumber // var can be changed

func main() {
	fmt.Println("starting-------------------")

	// variable
	var text string = "helloworld"
	var message = "hello hello" // assumes as string
	// Syntactic Sugar
	no_var_declare := 50
	fmt.Printf(text, message, no_var_declare)

	// integer
	var number int = 1
	fmt.Println(message)
	fmt.Println(number)

	// array
	var testarray = [50]string{"hi", "annyoung"}
	fmt.Printf("this is the array: %s\n", testarray)
	fmt.Printf("this it the test array %v and the size is %v\n", testarray[0], len(testarray))

	// slices
	var testslice = []string{"ashley", "baby", "christina"}
	testslice = append(testslice, "justin")
	fmt.Printf("this is the array: %v\n", testslice)
	fmt.Printf("this is the slice %v and the length is %v\n", testslice[2], len(testslice))

	//***** function execution exercise
	test_function(5)

	// fmt.Printf("value is %v \n", kiosk_function(remaining_tickets))
	fmt.Printf("value is %v \n", Kiosk_function_as_struct(remaining_tickets))

	// article.Article()

}
