package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "hello there friends!"
	// fmt.Println(strings.Contains(greeting, "hi"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "f"))
	// fmt.Println(strings.Split(greeting, " "))
	ages := []int{45, 32, 34, 2, 34, 2, 123, 1, 124, 752}

	sort.Ints(ages)

	index := sort.SearchInts(ages, 32)
	fmt.Println((index))

	names := []string{"yoshi", "mario", "peach", "browser", "luigi"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "browser"))
}
