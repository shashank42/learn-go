package chapters

import "fmt"

// Primitives : This is exposed
func Primitives() {
	fmt.Println("+++++++++++++++++PRIMITIVES+++++++++++++++++++++")
	var n bool = false // bool is its own datatype and not an alias for int
	fmt.Printf("%v, %T\n", n, n)

	m := 1 == 1
	o := 1 == 2
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", o, o)

	// Zero value of the numeric type sis 0
	// Integer : int is undefined size depending on the system being run on
	// int8, int16, int32, int64, further big package from math library
	// uint8, uint16, uint32, byte (8 bit unsigned integer)

	// Arithematic operations + - * / %

	// All operands in an expression should be the same type

	// Bit operators : & | ^(NOR) &^(AND NOT)
	// Bit shifts : << >>

	// Floating types
	// float32, float64 (default)
	// Bitwise, bitshift is not available

	// Complex numbers are built in
	// complex64, complex128
	var p complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", p, p)

	q := real(p)
	r := imag(p)
	fmt.Printf("%v, %T, %v, %T\n", q, q, r, r)

	s := complex(5, 12)
	fmt.Printf("%v, %T\n", s, s)

	// String are immutable
	t := "this is a string"
	fmt.Println(string(t[2]), t[3:10])
	// t[4] is actually of type uint8

	// String concatenation. Simply +
	// Convert string to slice of bytes/uint8
	b := []byte(t)
	fmt.Printf("%v, %T\n", b, b)

	// Why bytes - cause string is represented in system as files
	// Useful for sending data to different services on system

	// Rune : Different than string type
	// UTF8 is string <- String can be converted into collection of bytes/uint8
	// Rune is utf32 <- runes are true type alias of int32
	c := 'a'
	fmt.Printf("%v, %T\n", c, c)
	var d rune = 'b'
	fmt.Printf("%v, %T\n", d, d)

	fmt.Println("-----------------PRIMITIVES---------------------")
}
