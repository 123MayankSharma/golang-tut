package main

import "fmt"


func main(){
	//strings
	var name string = "Mayank"
	//in this type of declaration it is required to initialise at the time of declaration
	var name2 = "John"
	var name3 string
	fmt.Println(name,name2,name3);
	name="Ichi"
	name2="Cena"
	name3="initialised!"

	//this method of initialisation can't be used outside of a function
	//so to initialise a variable outside a function you need to use the above three methods.
        name4:="Type4"
	fmt.Println(name,name2,name3,name4);


	//ints
	 var age1 int =20
	var age2=30
	var age3 int 
	age3=40
	age4:=50
	fmt.Println(age1,age2,age3,age4)

	//variation of ints
	/*
		int8:[-128,127]
	*/
	var numO int8 = 25
	
	/*
          int16:[ -32768 through 32767]
we also have int32 and int64
	*/
	var numT int16 = 400

	//only positive numbers
	var numU uint = 100

	fmt.Println(numO,numT,numU)

	//floats
	var scoreOne float32 = 1.090920
        var scoreTwo float64 = -9.2191892192
	//in this method of declaration, by Default float64 is used due to higher precision
	scoreThree:=2.2101920190101290219021021091020192102910910910
	fmt.Println(scoreOne,scoreTwo,scoreThree)
	
}
