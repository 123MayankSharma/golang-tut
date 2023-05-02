package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a char: ")
	ch := ""
	ch, _ = reader.ReadString('\n')

	switch ch {
	case "a":
		fmt.Println("a....")
	case "b":
		fmt.Println("b....")
	default:
		fmt.Println("NOTA")

	}
}
