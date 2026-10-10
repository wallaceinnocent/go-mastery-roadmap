package functions

import "fmt"

func seperate(msg string) {
	fmt.Println(LineSeparator)
	fmt.Println("\t", msg)
	fmt.Println()
}
func Functions() {
	var greeting2 string
	greeting2 = "This is second greeting"

	sayHello()
	printStringFromArgument(greeting2)
	fmt.Printf("This value is returned from a function: %v, with formating\n", returnAvalueFromArgument(greeting2))
	fmt.Println("This value is returned from a function: ", returnAvalueFromArgument(greeting2))

	seperate("Functions basics")
	var age int
	age = 2

	fmt.Println("Original age: ", age)
	fmt.Println("Age in function that accepts age value: ", acceptValueArg(age))
	fmt.Println("Age after passing by value to a function: ", age)
	fmt.Println("Age in function that accepts age ref: ", *acceptRefArg(&age))
	fmt.Println("Age after passing by ref to a function and modified: ", age)

	// implementation of closures to remember section title
	runAnonymous := sectionMiddleware("Anonymous function")
	runHigerOrder := sectionMiddleware("Higher order functions")
	runClosures := sectionMiddleware("Closures")
	runCallbacks := sectionMiddleware("Callbacks")
	runNamedReturn := sectionMiddleware("Named return")
	runAnonymous(func() {
		fmt.Println(add(10, 30))
	})

	runClosures(func() {
		next := sequenceGenerator()
		fmt.Println(next())
		fmt.Println(next())
		fmt.Println(next())
	})

	runHigerOrder(func() {
		sum := doSomething(1, 5, func(i1, i2 int) int {
			return i1 + i2
		})
		fmt.Println("result from higher order function: ", sum)
	})
	runCallbacks(func() {

		processItems([]int{1, 2, 3, 4}, func(i int) {
			fmt.Printf("%d items processed\n", i)
		})
		fmt.Println()
		processItems([]int{1, 2, 3, 4, 3, 10, 4, 20}, funcCallback)
	})

	runNamedReturn(func() {
		s, p := doMath(10, 20)
		fmt.Printf("Sum is %d and Product is %d \n", s, p)
	})
}
func funcCallback(i int) {
	fmt.Printf("%d items processed\n", i)
}
