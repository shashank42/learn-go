package chapters

import "fmt"

const (
	_ = iota
	d
	e
	f
)

// Constants : Exported function
func Constants() {
	fmt.Println("+++++++++++++++++CONSTANTS+++++++++++++++++++++")
	const myConst int = 20 // math.Sin(30) won't work
	fmt.Printf("%v, %T\n", myConst, myConst)

	// Arrays can't be constants
	// Constants can be shadowed by constants or variables and any type

	const a = 42
	var b int16 = 27
	fmt.Printf("%v, %T\n", a+b, a+b) // This is allowed as GO actually replaces a with the literal - 42

	fmt.Println(d, e, f)

	fmt.Println("-----------------CONSTANTS---------------------")
}
