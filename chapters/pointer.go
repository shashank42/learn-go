package chapters

import (
	"fmt"
	//"unsafe"
)

// Pointers : Exposed
func Pointers() {
	// create pointers

	a := 33
	b := a // Copy by value
	fmt.Println(a, b)

	var c *int = &a
	// Asterisk here means it is declaring a pointer to the data type
	// & here gives the location of the data element
	// Dereferencing pointers
	a = 343
	fmt.Println(a, b, c, *c) // Asterisk here means it is dereferencing the pointer variable

	*c = 44
	fmt.Println(a, b, c, *c) // Both values changed again

	// Pointer arithematic - Go does not allow such math allowed
	d := [3]int{1, 2, 3}
	e := &d[0]
	f := &d[1]
	fmt.Printf("%v, %v %v, %p, %p\n", d, e, f, e, f)
	*e = 45
	fmt.Printf("%v, %v %v, %p, %p\n", d, e, f, e, f)

	// If you still want pointer arithematic you can use unsafe package

	type myStruct struct {
		foo int
	}
	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	var ms1 *myStruct
	fmt.Println(ms1) // This will print nil and that is what you need
	ms1 = new(myStruct)
	fmt.Println(ms1)
	(*ms1).foo = 342
	// But this also works
	ms1.foo = 456 // Syntactic sugar. Only for cleanness.
	fmt.Println(*ms1, ms1, (*ms1).foo, ms1.foo)

	g := [3]int{1, 3, 5}
	h := g // Array is copied by value

	fmt.Println(g, h)
	h[1] = 43
	fmt.Println(g, h)

	i := []int{1, 3, 5}
	j := i // Slice is copied by reference

	fmt.Println(i, j)
	j[1] = 43
	fmt.Println(i, j)

	// Maps also get passed by reference

	// The new function

	// Working of nil

	// Types with internal pointers
}
