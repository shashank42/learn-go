package chapters

import "fmt"

// Looping : Exported
func Looping() {
	fmt.Println("+++++++++++++++++LOOPING+++++++++++++++++++++")

	// Simple loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 5; i = i + 2 { // Different incrementor
		fmt.Println(i)
	}

	for i, j := 0, 1; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// Exit a loop early
	j := 0
	for {
		fmt.Println(j)
		j++
		if j == 5 {
			break
		}
	}
	// Use of Labels
Loop:
	for k := 1; k <= 3; k++ {
		for l := 1; l <= 3; l++ {
			fmt.Println(k * l)
			if k*l >= 4 {
				break Loop
			}
		}
	}

	// Looping over collections

	// Loop over slice/array
	s := []int{1, 2, 3}
	fmt.Println(s)
	for k, v := range s {
		fmt.Println(k, v)
	}

	// Loop over map
	statePopulation := map[string]int{
		"California": 231232,
		"Texas":      343423432,
		"Florida":    34324,
	}
	fmt.Println(s)
	for k, v := range statePopulation {
		fmt.Println(k, v)
	}

	// Loop over string
	st := "Hello go!"
	for k, v := range st {
		fmt.Println(k, v, string(v))
	}

	for _, v := range st {
		fmt.Println(v, string(v))
	}

	// Can also range over channel

	fmt.Println("-----------------LOOPING---------------------")
}
