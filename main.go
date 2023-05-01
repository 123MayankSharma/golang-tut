package main

import "fmt"

// basic demonstration of function
func sumOfSlice(mySlice []int) int {
	sum := 0
	mySlice[0] = 5
	for _, value := range mySlice {
		sum += value
	}

	return sum
}

func intro(s string) {
	fmt.Println("Word Length is:", len(s))
}

// take function as argument
func callbackFunc(dict []string, f func(string)) {
	for _, val := range dict {
		f(val)
	}
}

// how to return multiple values
func multipleReturnValues(dict []string) (int, int) {
	countOfString := len(dict)
	totalChars := 0
	for _, val := range dict {
		totalChars += len(val)
	}

	return countOfString, totalChars
}
func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	a := sumOfSlice(mySlice)
	fmt.Println("Sum of ", mySlice, "is", a)
	callbackFunc([]string{"my", "name", "is", "mayank"}, intro)
	cs, tc := multipleReturnValues([]string{"my", "name", "is", "mayank"})

	fmt.Println(cs, tc)
}
