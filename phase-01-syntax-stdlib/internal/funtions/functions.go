package functions

import "fmt"

func sayHello() {

	fmt.Println("This is hello from a function")
}
func printStringFromArgument(message string) {

	fmt.Println(message)
}

func returnAvalueFromArgument(input any) any {

	return input
}

func acceptValueArg(age int) int {
	return age + 10
}
func acceptRefArg(age *int) *int {
	*age = 10
	return age

}
