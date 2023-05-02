package externalFile

import (
	"fmt"
)

// methods and variables with only capital names are visible outside of the package
// but inside the package there is no such convention
func printName() {
	fmt.Println("Maank")
}
