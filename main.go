package main

import "fmt"

func main() {
	//illustrated below is golang while loop
	i := 0
	for i < 5 {
		fmt.Println("Iteration: ", i)
		i++
	}

	/*Illustrated below is golang for loop*/
	for i := 0; i < 10; i++ {
		fmt.Println("For-Loop Iter: ", i)
	}

	//iterating through an array
	mySlice := []string{"Hey", "How", "are", "you!!!"}

	for i := 0; i < len(mySlice); i++ {
		fmt.Println("mySlice element at idx ", i, "is ", mySlice[i])
	}

	/*
	iterating through array using for loop with range
	here value is like a local value so even if we change value like val=abc the original slice is not affected*/
	for idx, val := range mySlice {
		fmt.Println("mySlice element at idx ", idx, "is ", val)
	}

	//if you only need value/index in for loop replace the other one with _
	for _, val := range mySlice {
		fmt.Println("value: ", val)
	}

}
