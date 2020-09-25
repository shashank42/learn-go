package chapters

import "fmt"
import "strconv"

// Declare variables at package level
var n int = 345

// N : Upper case variables are visible at package level. Rest are not.
// Need not be compulsorily used
var N int = 343

// o := 34 Not allowed

// Variable block at application level
var (
	actorName    string = "Shashank Srivastava"
	directorName string = "Akku Dua"
	season       int    = 3
)

// Variables : This is the exposed function
func Variables() {
	fmt.Println("+++++++++++++++++VARIABLES+++++++++++++++++++++")
	fmt.Println(42)

	var i int      // As you read it. Declare but not initialize
	fmt.Println(i) // Initializes to default value 0
	i = 42
	fmt.Println(i)

	var j int = 42
	fmt.Println(j)

	k := 42
	fmt.Println(k)

	// Error : no new variables on left side of :=
	// i := 42

	fmt.Printf("%v, %T \n", j, j)

	var l float64
	fmt.Printf("%v, %T \n", l, l)

	m := 45.34
	fmt.Printf("%v, %T \n", m, m)

	// Print the package level n
	fmt.Printf("%v, %T \n", n, n)

	n := 23.56 // Create the shadowing variable of the package level variable
	fmt.Printf("%v, %T \n", n, n)

	// Variables always have to be used

	// Type cast should be explicit
	var o float32
	o = float32(n)
	fmt.Printf("%v, %T \n", o, o)

	// Convert values to string is a pain -> cannot convert o (type float32) to type string
	// var st string
	// st = string(o)
	// fmt.Printf("%v, %T \n", st, st)

	var st string
	st = strconv.Itoa(k)
	fmt.Printf("%v, %T \n", st, st)

	// For floats use Sprintf
	st = fmt.Sprintf("%f", o)
	fmt.Printf("%v, %T \n", st, st)

	// Summary				      .							    .
	fmt.Println("-----------------VARIABLES---------------------")

}
