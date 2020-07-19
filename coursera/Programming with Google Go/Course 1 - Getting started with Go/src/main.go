package main

import "fmt"

// rename a variable type with a custom value
type UUID string

// short declaration and assignations do not work out of methods (global scope)
// notValidVar := "invalid"

func main() {
	fmt.Printf("Hello, world!\n")
	initVariables()
}

func initVariables() {
	// declare variable and assign a value to it at the same time
	var intVar int = 10
	// declare variable and assign a value later
	var intVar2 int
	intVar2 = 100
	// use custom type to declare and assign a value to a variable
	var uuidVar UUID = "This is an id"
	// short declaration and assignation of a int var, it only works inside methods
	shortVar := 5
	// it is also valid to declare a Variable without specifying the type, go compiler will infer it
	var nonTypeVar = "infered value as string"
	var sumVars = intVar + intVar2 + shortVar
	fmt.Printf("Sum of int variables is %d\n", sumVars)
	fmt.Printf("uuidVar: %s and nonTypeVar: %s\n", uuidVar, nonTypeVar)
}
