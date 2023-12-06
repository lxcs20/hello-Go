package main

import (
	"fmt"
)

func main() {

	fmt.Println("hello! Go")
	fmt.Println("Go! Go!")

	// struct_()
	// variable()
	// ifControl()
	// switchCase()
	// shortIf()
	// forLoop()
	// doWhileLoop()
	whileLoop()
}

func variable() {
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
	fname    string
	lname    string
	birthday string
	gender   string
}

func struct_() {
	person := Person{
		fname:    "lanthong",
		lname:    "myLast",
		birthday: "0X-0X-X00X",
		gender:   "M",
	}

	fmt.Println("hello, ", person)
}

func ifControl() {
	num := 10
	if num >= 10 {
		fmt.Println(num)
	} else {
		fmt.Println("else case")
	}

	if num < 10 {
		fmt.Println(num)
	} else if num == 5 {
		fmt.Println(num)
	} else {
		fmt.Println(num)
	}
}

func switchCase() {
	str := "b"
	switch str {
	case "a":
		fmt.Println(str)
	case "b":
		fmt.Println(str)
	case "bb":
		fmt.Println(str)
	default:
		fmt.Println("not matching")
	}
}

func shortIf(){
	num1 := 5
	num2 := 10

	if sumNum := num1 + num2; sumNum >= 10 {
		fmt.Println(sumNum)
	} 
}

func forLoop(){
	for i:=0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func doWhileLoop (){
	i:=1
	for {
		fmt.Printf("number: %d\n", i)
		i++
		if(i >= 10){
			break
		}
	}	
}

func whileLoop (){
	i:=0
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}

