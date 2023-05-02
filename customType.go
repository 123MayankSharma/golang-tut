package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// this is how structs are defined in go
// we can also add methods to structs and they are known as receiver functions
// and they are equivalent of methods of a class
type customType struct {
	name  string
	items map[string]float64
}

func customTypeObjectGenerator(name string, itemNames []string, itemCost []float64) customType {
	objectMap := map[string]float64{}
	for idx, val := range itemNames {
		objectMap[val] = itemCost[idx]
	}
	myObject := customType{
		name:  name,
		items: objectMap,
	}

	return myObject
}

// here we have added (ct customType) before the function name which
// basically means that this function can only be called by a custom type object essentially becoming a method
func (ct customType) printMap() {
	for key, val := range ct.items {
		fmt.Println(key, "-", val)
	}
}

//we can also have receiver function with pointers
//given below is an example

func takeInputOfRandomString(randomString *string, r *bufio.Reader) {
	fmt.Println("Enter value of random string")

	*randomString, _ = (*r).ReadString('\n')
	*randomString = strings.TrimSpace(*randomString)
	fmt.Println("Random String: ", *randomString)
}
func (ct *customType) updateItemPrice(item string) {
	//we access the members of struct in case of a pointers
	//by using the below syntax: (*ct).xyz_member_of_struct
	// and we will be taking input from user
	price := 0.0
	randomString := ""
	//below 4 lines are used to take line as input
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Type of bufio.NewReader is %T \n", reader)
	takeInputOfRandomString(&randomString, reader)
	fmt.Println("Enter the new price for ", item)
	fmt.Scan(&price)
	(*ct).items[item] = price
}
