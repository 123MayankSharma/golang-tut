package main

import "fmt"

func main() {
	//here string is key type and int is value type
	mp := map[string]int{"Apple": 0}

	//assigning values in map
	mp["Mayank"] = 1
	mp["sharma"] = 2

	//printing value from a map
	fmt.Println(mp["Apple"])

	//looping over a map
	for key, value := range mp {
		fmt.Println(key, ": ", value)
	}

}
