package main

import "fmt"

func main() {
	a := 5
	if a < 5 {
		fmt.Println("Lesser")
	} else if a == 5 {
		fmt.Println("equal")
	} else if a > 5 {
		fmt.Println("Greater")
	}
}
