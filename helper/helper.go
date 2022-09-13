package helper

import (
	"fmt"
)

func Check_valid_city(city string) bool {
	//switch
	switch city {
	case "London", "Frankfurt":
		fmt.Printf("city is in Europe : %s\n", city)
		return true
	case "Singapore", "HongKong", "Seoul":
		fmt.Printf("city is in asia\n")
		return true
	default:
		fmt.Printf("city doesn't belong in both ...\n")
		return false
	}

}
