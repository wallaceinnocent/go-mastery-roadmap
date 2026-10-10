package functions

/*

sectionMiddleware is a Higher-Order Function (The Factory): It takes a parameter (title string) and returns a function.

The Returned Function is a Closure: Because the inner anonymous function references title, it captures and remembers that specific string in its lexical scope. Even after sectionMiddleware finishes running and exits, title stays alive in memory for that specific returned instance.

The Closure is also a Higher-Order Function: It accepts a function (task func()) as its argument.

The Execution Flow: When you invoke it, it combines all three things:

It uses the remembered title to call seperate(title).

It immediately executes the passed-in task() right after.
*/
func sectionMiddleware(title string) func(func()) {
	return func(task func()) {
		seperate(title)
		task()
	}
}
