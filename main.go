package main

import (
	"fmt"

	structpackage "github.com/lxcs20/hello-go/structPackage"
)

func main() {

	fmt.Println("hello! Go")
	fmt.Println("Go! Go!")

	var student1 structpackage.Student
	student1 = structpackage.Student{}
	student1.Setfname("lanthong")
	student1.Setlname("xyt")
	student1.Setage(20)

	fmt.Println(student1.Getfname())

	// variable()
	// ifControl()
	// switchCase()
	// shortIf()
	// forLoop()
	// doWhileLoop()
	// whileLoop()
	// Array()
	// slice()
	// Map()
	// struct_()
}

func ariable() {
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

func shortIf() {
	num1 := 5
	num2 := 10

	if sumNum := num1 + num2; sumNum >= 10 {
		fmt.Println(sumNum)
	}
}

func forLoop() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func doWhileLoop() {
	i := 1
	for {
		fmt.Printf("number: %d\n", i)
		i++
		if i >= 10 {
			break
		}
	}
}

func whileLoop() {
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}

func Array() {
	var Arr [3]int
	Arr[0] = 1
	Arr[1] = 2
	Arr[2] = 3
	fmt.Println(Arr)
}

func slice() {
	var slice []int
	slice = append(slice, 1)
	slice = append(slice, 2, 3, 4)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(slice[0:1])
	fmt.Println(slice)

}

func Map() {
	myMap := make(map[string]int)

	myMap["apple"] = 3
	myMap["banana"] = 5
	myMap["orage"] = 9

	fmt.Println("Apple: ", myMap["apple"])
	myMap["banana"] = 19

	delete(myMap, "apple")
	for key, value := range myMap {
		fmt.Printf("%s -> %d\n", key, value)
	}

	val, ok := myMap["melon"]
	if ok {
		fmt.Println("melon is: ", val)
	} else {
		fmt.Println("melon not found in myMap")
	}

	valOr, isHave := myMap["orage"]
	if isHave {
		fmt.Println("orange is ", valOr)
	} else {
		fmt.Println("orage not found in myMap")
	}
	fmt.Println(valOr)
}

type Person struct {
	fname    string
	lname    string
	birthday string
	gender   string
	address Address
}

type Address struct {
	village  string
	district string
	province string
}

func struct_() {
	myInfo := Person{
		fname:    "lanthong",
		lname:    "myLast",
		birthday: "0X-0X-X00X",
		gender:   "M",
		address: Address{
			village: "vill",
			district: "dist",
			province: "province",
		},
	}

	fmt.Println("hello, ", myInfo)

	var studentS [2]Person

	studentS[0] = Person{
		fname: "std1",
		lname: "std2",
		birthday: "0X-0X-20XX",
		address: Address{
			village: "std1 vill",
			district: "std1 dist",
			province: "std1 prov",
		},
	}

	fmt.Println(studentS[0])
}
