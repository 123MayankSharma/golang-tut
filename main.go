package main

import "fmt"

func main() {
	//this is how object of customTypes are generated
	customObject := customTypeObjectGenerator("Mayank", []string{"Onion", "Potato"}, []float64{10.5, 11.2})
	fmt.Println(customObject)
	customObject.printMap()
	(&customObject).updateItemPrice("Onion")
	fmt.Println(customObject)

	//now we will be taking input from user

}
