package main

import "fmt"


func main(){
	//arrays and multiple ways of 
	//declaring them
	var ages [5]int =[5]int{20,30,40,50,60}	
	fmt.Println(ages)
	var ages2=[6]int{1,2,3,4,5,6}
        ages4:=[10]int{1,2,3,4,5,6,7,8,9,10} 
	fmt.Println(ages2,ages4)
        
	//slices and multiple ways of declaring them
	var ages3=[]int{1,2,3,4,5,6,7,8,9}
	//method to append element to slice
	ages3 = append(ages3,999)
	var ages5 []int = []int{5,10,15,20}
	fmt.Println(ages5)
	var ages7 []int
	ages7=[]int{7,14,21,28}
	fmt.Println(ages7)
	ages9:=[]int{9,18,27,36,45}
	fmt.Println(ages9)


	//method to print length of array/slice
	fmt.Println(len(ages5))

	//get item at a specific index
	fmt.Println(ages[2])

	/*
	slice ranges allow us to get a range of a slice and return it as a new slice.
here last index is excluded
	*/

	rangeOfAges5:=ages5[0:2]
	fmt.Println("Range 0:2 to ages5 is ",rangeOfAges5)

	//how to append one slice with another slice
	ages5=append(ages5, rangeOfAges5...)
	fmt.Println(ages5)
}
