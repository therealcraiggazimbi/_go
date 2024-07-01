package main

import "fmt"

func main() {

	//Define map
	// emails := make(map[string]string)

	//Assign kv
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	emails := map[string]string{"Bob": "bob@gmail.com"}
	fmt.Println(emails)

	//delete
	delete(emails, "Bob")

}
