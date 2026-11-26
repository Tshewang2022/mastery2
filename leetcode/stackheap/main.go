package main

import "fmt"

type User struct {
	name string
}

func main() {
	user := newUser("Alice")
	fmt.Println(user.name)
}

func newUser(name string) *User {
	user := &User{name: name}
	return user
}

// stack vs heap from christopher Berger
// main questions=> how does the some variable does get escaped to heap

// in some language, stack has the upper bound, but in go, if its run out of space, it allocate more
// space,
// WORKING MECHANISM => LIFO
// main function=> stores local variable and functions

//*** when is value possibly constructed on the heap?
// 1. when a value could possibly be referenced after the function that constructed the value returns
// 2. when the compiler determins a value is too large to fit on the stack
// 3. when the compiler doesnt know the size of a value at the compile time

// compile vs run time

// some commonly allocated values
// values shared with pointers
// variable stored in interface variables
// func literal variables => variable captured by a closure

// backing data for maps, channels, slices and strings

// points to remember
// optimize for correctess, not performance
// go only put functions variables on the stack if it can prove a varible is not used after the function returns
// sharing down typically escapes to the heap
// sharing up typically escapes to the heap
// asking the compiler to find out
// dont guess. User the tooling

// wow, such a insights in a small video
