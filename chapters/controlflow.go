package chapters

import "fmt"

// ControlFlow : Control flow exposed
func ControlFlow() {
	fmt.Println("+++++++++++++++++ControlFlow+++++++++++++++++++++")

	if true {
		fmt.Println("The test is true")
	}

	statePop := map[string]int{
		"Cali": 34543345,
		"Tex":  3476484,
		"Res":  235237644,
	}
	if pop, ok := statePop["Res"]; ok {
		fmt.Println(pop, ok)
	}

	number := 50
	guess := 60
	if guess < number {
		fmt.Println("Too low")
	} else {
		fmt.Println("Too high")
	}

	switch 3 {
	case 1:
		fmt.Println("case 1")
	case 2, 3, 4:
		fmt.Println("case 2")
	default:
		fmt.Println("not one or two")
	}

	// Default behavior
	switch 2 {
	case 1:
		fmt.Println("case 1")
	case 2, 3, 4:
		fmt.Println("case 2")
	default:
		fmt.Println("not one or two")
	}

	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than 10")
		fallthrough // You have to manually manage when there is a fallthrough
		// The default behaviour is "break"
	case i <= 20:
		fmt.Println("less than 20")
	default:
		fmt.Println("greater than twenty")
	}

	// var f interface{} = 1
	var f interface{} = [3]int{1, 2, 3}
	switch f.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is float64")
	case string:
		fmt.Println("i is string")
	case [2]int: // The type or the array is the type AS WELL AS the size of the array
		fmt.Println("i is [2]int")
	default:
		fmt.Println("i is another type")
	}

	fmt.Println("-----------------ControlFlow---------------------")
}
