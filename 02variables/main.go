package main

import "fmt"

const AccessToken string = "ejJh..." // Capital variable names refer that it is Public specified

// primitive data types
// string

// bool

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32
//      // represents a Unicode code point

// float32 float64

// complex64 complex128

func main() {
	// var type variable_name = value_assigned
	var name string = "foreverabhi"
	fmt.Println("name: ", name) // name: foreverabhi
	fmt.Printf("%T \n", name)   // string

	var isBooleanVal bool = true
	fmt.Println("isBooleanVal: ", isBooleanVal) // isBooleanVal: true
	fmt.Printf("%T \n", isBooleanVal)           // bool

	var smallVal = 555
	fmt.Println("smallVal: ", smallVal) // smallVal: 555
	fmt.Printf("%T \n", smallVal)       // int

	var smallFloat float32 = 555.46678
	fmt.Println("smallFloat: ", smallFloat) // smallVal: 555.4668
	fmt.Printf("%T \n", smallFloat)         // float32

	// default value and some aliases
	var anotherVariable int
	fmt.Println("anotherVariable: ", anotherVariable) // anotherVariable: 0
	fmt.Printf("%T \n", anotherVariable)              // int

	// implicit type
	var notDeclaringType = "Gotcha!"
	fmt.Println("notDeclaringType: ", notDeclaringType) // notDeclaringType: Gotcha!
	fmt.Printf("%T \n", notDeclaringType)               // string

	// no var style
	noVarStyle := "Creakkk"
	fmt.Println("noVarStyle: ", noVarStyle) // noVarStyle: Creakkk
	fmt.Printf("%T \n", noVarStyle)         // string

	fmt.Println("Public Access Token: ", AccessToken)
}
