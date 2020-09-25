package chapters

import "fmt"

// Functions : Exported
func Functions() {
	fmt.Println("+++++++++++++++++Functions+++++++++++++++++++++")
	// Basic syntax

	// Parameters
	printString("printString function", "shashank")

	string1 := "Raj"
	string2 := "Kumar"
	printRefString(&string1, string2)
	fmt.Println(string1, string2)

	// Return values

	s := bariaticParameter(1, 2, 3, 4, 5) // This is an example of the bariatic arguments also
	fmt.Println("The sum is ", *s)
	// Lot of things are happening - it's passing variables as reference as well
	// Bariatic, pointers as return types

	s1 := namedReturnValue(1, 2, 3, 4, 5) // Not often used
	fmt.Println("The sum is ", s1)

	d, err := divide(1.0, 0.0)
	if err != nil {
		fmt.Println(err)
		//return
	} else {
		fmt.Println("Division answer is : ", d)
	}

	// Anonymous functions
	func() {
		if 1 == 2 {
			fmt.Println("Same")
		} else {
			fmt.Println("Different")
		}
	}()

	// Functions as types
	// Can even use var f func() = func(){ some statements }
	f := func() {
		if 1 == 2 {
			fmt.Println("Same")
		} else {
			fmt.Println("Different")
		}
	}
	f()
	fmt.Printf("%T, %v\n", f, f)

	// Methods
	g := greeter{
		greeting: "Hello",
		name:     "Shashank",
	}
	g.greet()

	fmt.Println("-----------------Functions---------------------")
}

func printString(msg, person string) { // Argument // Syntactic sugar
	fmt.Println(msg, person)
}

func printRefString(string1 *string, string2 string) { // Argument // Syntactic sugar
	fmt.Println(*string1, string2)
	*string1 = "Rajendra" // Only this will get changed
	string2 = "Kukur"
}

func bariaticParameter(values ...int) *int { // Says that take the last arguments and wrap them up into a slice of type int
	fmt.Printf("%v, %T\n", values, values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("Sum of the values is : ", result)
	return &result
}

func namedReturnValue(values ...int) (result int) { // Says that take the last arguments and wrap them up into a slice of type int

	fmt.Printf("%v, %T\n", values, values)
	for _, v := range values {
		result += v
	}
	fmt.Println("Sum of the values is : ", result)
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() { // g is passed as value. For call by reference you can call with (g *greeter)
	fmt.Println(g.greeting, g.name)
}
