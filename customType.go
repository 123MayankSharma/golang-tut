package main

import "fmt"

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

func (ct *customType) updateItemPrice(item string, price float64) {
	//we access the members of struct in case of a pointers
	//by using the below syntax: (*ct).xyz_member_of_struct
	(*ct).items[item] = price
}
