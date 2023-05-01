package main

import "fmt"


func main(){
	//Print does not include newline
	//whereas Println does include newline
	fmt.Print("Hello \n")
	fmt.Print("World")


	/*Printf formatted strings*/
	fmt.Printf("My name is %v and my age is %v \n","Mayank",22)
	/*used to print given variable with quotes(q). works with strings*/
	fmt.Printf("My name is %q and my age is %q \n","Mayank",22)
	/*To print type of variable*/
	age:=22
	fmt.Printf("type of age is %T \n",age)



	//another way of formatting strings
	//with type
	//for float where in 0.Xf X represents to how much decimal number float is rounded
	fmt.Printf("My age is %0.1f \n",22.12)	


	//for integer
	fmt.Printf("My age is %d \n",22)	

	//for string
	fmt.Printf("My name is %s","Goku \n")	

	/*

!!! Sprintf

We can use it to save the formatted string to a variable
	*/
	formatted_string := fmt.Sprintf("My age is %v and my name is %v",20,"Spiderman")
	fmt.Println(formatted_string)

}
