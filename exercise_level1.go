package main

import "fmt"

// Exercise 2
var a int
var b string
var c bool

// Exercise 4
type droid int
var d droid

//Exercise 5
var e int

func main() {
	/* 
	Exercise 1
	1.Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
		a. 42
		b. “Resistance is futile!”
		c. true
	2. Now print the values stored in those variables using ...
		a. a single pritn statement
		b. multiple print statements
	*/

	x := 42
	y := "Resistance is futile!"
	z := true

	fmt.Printf("%v\t%v\t%v\n", x, y, z)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	/* 
	Exercise 2
	1. Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
		a. identifier “a” type int
		b. identifier “b” type string
		c. identifier “c” type bool
	2. in func main
		a. print out the values for each identifier
		b. The compiler assigned values to the variables. What are these values called?
	*/

	fmt.Printf("%v\t%v\t%v\n", a, b, c)

	// The compiler assigned values are called zero values.

	/*
	Exercise 3
	Using the code from the previous exercise,
	1. At the package level scope, assign the following values to the three variables
		a. for a assign 42
		b. for b assign “James Bond”
		c. for c assign true
	2. in func main
		a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
		b. print out the value stored by variable “s”
	*/

	a = 42
	b = "James Bond"
	c = true

	s := fmt.Sprintf("%v\t%v\t%v\n", a, b, c)
	fmt.Println(s)

	/*
	Exercise 4
	1. Create your own type. Have the underlying type be an int.
	2. Create a VARIABLE of your new TYPE with the IDENTIFIER “d” using the “VAR” keyword
	3. In func main
		a. print out the value of the variable “d”
		b. print out the type of the variable “d”
		c. assign 42 to the VARIABLE “d” using the “=” OPERATOR
		d. print out the value of the variable “d”
	*/

	fmt.Println(d)
	fmt.Printf("%T\n", d)
	d = 42
	fmt.Println(d)

	/*
	Exercise 5
	Building on the code from the previous example
	1. At the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER "e". The variable should be of the UNDERLYING TYPE of your custom TYPE.
	2. In func main use CONVERSION to convert the TYPE of the VALUE stored in “d” to the UNDERLYING TYPE
		a. then use the “=” operator to ASSIGN that value to “e”
		b. print out the value stored in “e” 
		c. print out the type of “e”
	*/

	e = int(d)
	fmt.Println(e)
	fmt.Printf("%T\n", e)
}