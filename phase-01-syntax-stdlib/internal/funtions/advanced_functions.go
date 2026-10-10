package functions

import "fmt"

/*
  1. Anonymous Functions (Function Literals)
 These are functions declared without a name.
 They can be assigned to variables, passed into other functions,
  or executed immediately (Immediately Invoked Function Expressions - IIFE)
*/
var add = func(a int, b int) int {
	return a + b
}

/*
 2. Closures
 A closure is a special type of anonymous function that references variables from outside its body.
 The function can access and assign to the referenced variables even after the outer function has finished executing.
*/

func sequenceGenerator() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

/*
Independent State (Multiple Instances)
If you create a second generator, it gets its own completely separate, isolated instance of i:

seqA := sequenceGenerator()
seqB := sequenceGenerator()

fmt.Println(seqA()) // Output: 1
fmt.Println(seqA()) // Output: 2

fmt.Println(seqB()) // Output:
*/

// The use of closures

/*
1. State Encapsulation (Data Hiding)
In object-oriented languages, you use private class fields to hide state. In Go, you can use closures to maintain private state without needing a full struct. Because i lives entirely inside the closure, nothing outside of next() can accidentally modify or tamper with i.

2. Generator Patterns & Iterators
As shown in your example, closures are great for generating unique sequences, auto-incrementing IDs, or custom iterators where each call needs to pick up right where the last one left off.

3. Middleware and Decorators (Functional Programming)
In web servers (like standard Go HTTP handlers), closures are heavily used to "wrap" functions with extra behavior—such as logging, authentication, or timing execution—without changing the core function's code.

4. Callback Functions and Event Handling
Closures are frequently used in asynchronous programming or event listeners where an inner function needs access to variables available in the scope where it was registered.
*/

/*
 * 3. Higher-Order Functions
 A higher-order function is a function that takes one or more functions as arguments, returns a function, or both.
*/

func doSomething(a, b int, opperation func(int, int) int) int {
	return opperation(a, b)
}

// callback
//
/*
 * It is a function passed as an argument to another function, which is then invoked
 * ("called back") at a later point or when a specific condition is met
 */

func processItems(items []int, callback func(int)) {
	i := 0

	for _, t := range items {
		fmt.Printf("Item %d processed\n", t)
		i++
	}
	callback(i)
}

/*
 * You can name the return values in the function signature.
 *  Go will automatically initialize them as zero-values, and a bare return statement will return them.
 */

func doMath(a int, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return
}
