package main

import "fmt"

func main() {

	ids := []int{33, 76, 54, 23, 11, 2}

	//Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Loop through ids
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Range with map

	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

}
