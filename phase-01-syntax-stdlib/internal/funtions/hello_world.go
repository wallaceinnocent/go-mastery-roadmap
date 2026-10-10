package functions

import "fmt"

var LineSeparator = "--------------------------*--------------------------"

func HelloWorld() {
	greeting := "Hello world"

	var greeting2 string
	greeting2 = "This is second greeting"

	fmt.Println(greeting)
	fmt.Println(greeting2)

}
