package main

import (
	"fmt"
	"os"
)

func main() {
	mp := map[string]int{}

	mp["a"] = 1
	mp["b"] = 2
	mp["c"] = 3
	//fill with all english characters
	mp["d"] = 4
	mp["e"] = 5
	mp["f"] = 6

	mapString := ""
	for key, value := range mp {
		mapString += fmt.Sprintf("%s:%d\n", key, value)
	}
	//converting string to  byte array
	data := []byte(mapString)
	os.WriteFile("TestFile.txt", data, 0644)
}
