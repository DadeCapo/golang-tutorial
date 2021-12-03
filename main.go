package main

import "fmt"

func main() {

	//strings
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"

	fmt.Println(nameFour)

	//intergers

	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	//bits and memory

	var numOne int8 = 25
	var numTwo = -128
	var numThree uint16 = 256

	//floats

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 52456245252523532532.7

}
