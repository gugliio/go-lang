package main

import "fmt"

func namingConventions() {
	var mv int        // max value
	var max_value int // NOT OK
	fmt.Println(mv, max_value)

	var packetsReceived int //NOT OK, name to long
	var b int               //OK
	fmt.Println(b, packetsReceived)

	const MAX_VALUE = 100 // NOT OK
	const N = 100         // BETTER

	// Exported value
	var Max = 100
	// Private value
	var test = 100
	fmt.Println(Max, test)

	writeToDB := true  // ok, idiomatic
	writeToDb := false // not ok

	fmt.Println(writeToDB, writeToDb)
}
