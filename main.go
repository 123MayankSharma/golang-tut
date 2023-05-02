package main

import "fmt"

func changeString(s *string) {
	//we can access value by using *
	fmt.Println("Initial value of s is: ", *s)
	*s = "Mayank"
}

func main() {
	var s string = "Rahul"
	//we pass by reference by using & operator while calling
	//and variable *data_type while using
	// we access pointer values by using * as illustrated above
	changeString(&s)
	fmt.Println(s)
}
