package main

import (
	"fmt"
	"sort"
	"strings"
)


func main(){
	myGreeting:="Hey How are you"

	//string standard library method to check 
	//if a string contains a piece of text
	//for more info refer to strings module methods
	fmt.Println(strings.Contains(myGreeting,"How are"))


	/*
	strings.replace to replace all instances of one word/text with another. It does not modify original string.

	*/
	fmt.Println(strings.ReplaceAll(myGreeting,"How are you","where have you been"))


	/*
	 Convert String to upper case. It returns a new String and does not modify the original string
	*/
	fmt.Println(strings.ToUpper(myGreeting),myGreeting)

	/*
	strings.Index to find first Index of a text
	*/
	fmt.Println(strings.Index(myGreeting,"ar"))

	//just like javascript split
	fmt.Println(strings.Split(myGreeting," "))


	//sort package for Ints,Floats etc
	//and modifies the original package
	mySlice:=[]int{8,2,1,5,3,6}
	sort.Ints(mySlice)
	fmt.Println("Sorted mySlice is: ",mySlice)


	/*Search position of an element in an sorted slice/array and if element doesn't exist we are returned length of the array as index*/
	idx:=sort.SearchInts(mySlice,3)
	fmt.Println("position of 3 is ", idx)
	
	/*Sort strings*/
	myStringSlice:=[]string{"my","name","is","mayank"}
sort.Strings(myStringSlice)
	fmt.Println("Sorted myStringSlice is: ",myStringSlice)
	


}
