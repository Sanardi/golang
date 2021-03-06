package main

import (
	"fmt"

	"github.com/pluralsight/webservice/models"
)

func main_old() {

	arr := [3]string{}
	arr[0] = "apple"
	arr[1] = "plum"
	arr[2] = "orange"
	fmt.Println(arr)

	b := contains(arr, "plum")
	fmt.Println(b)

	b = contains(arr, "grape")
	fmt.Println(b)

	u := models.User{
		ID:        2,
		FirstName: "Marzia",
		LastName:  "Azam",
	}
	fmt.Println(u)
}

func contains(arr [3]string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
