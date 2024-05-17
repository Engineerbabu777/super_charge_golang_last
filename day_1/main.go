package main

import "fmt"

// DATA TYPES!

var (
	VarFloat32 float32 = 0.6
	VarFloat64 float64 = 0.6
	name       string  = "Hello"
	VarInt32   int32   = 06
	VarInt64   int64   = 06
)

type Name struct {
	Base string
}

func main() {

	// give me examples for byte data type in golang!
	var byteVar byte = 0x10
	fmt.Println(byteVar)

	users := map[string]int{
		"user1": 7,
	}

	fmt.Printf("%v\n", users)

	var myArray [2]string

	myArray = [2]string{"abc", "def"}

	fmt.Println(myArray)

	mySlices := []string{"abc", "def"}

	// myArray = append(myArray, "abc2") // throw an error!

	fmt.Println(myArray, mySlices)

}
