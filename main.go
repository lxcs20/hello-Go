package main

import (
	"fmt"
)


func main() {
	
	fmt.Println("hello! Go")
	fmt.Println("Go! Go!")
	
	// struct_()
	// variable()
	ifControl()
}

func variable () {
	var fname string
	var lname string
	var age int
	var float float32
	var boolean bool

	fname = "lanthong"
	lname = "myLast"
	age = 20
	float = 3.14
	boolean = true

	fmt.Println(fname)
	fmt.Println(lname)
	fmt.Println(age)
	fmt.Println(float)
	fmt.Println(boolean)
}

type Person struct {
	fname string
	lname string
	birthday string
	gender string
}
func struct_ () {
	person := Person{
		fname: "lanthong",
		lname: "myLast",
		birthday: "0X-0X-X00X",
		gender: "M",
	}

	fmt.Println("hello, ", person)
}

func ifControl (){
	num := 10
	if(num >= 10){
		fmt.Println(num)
	} else {
		fmt.Println("else case")
	}

	if(num < 10){
		fmt.Println(num)
	} else if (num == 5){
		fmt.Println(num)
	} else {
		fmt.Println(num)
	}
}


