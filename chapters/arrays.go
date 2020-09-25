package chapters

import "fmt"

// Arrays : Exported
func Arrays() {
	fmt.Println("+++++++++++++++++ARRAYS+++++++++++++++++++++")
	grades := [...]int32{97, 85, 99}
	var students [3]string
	fmt.Println(grades)
	fmt.Println(students)

	// Arrays are variables, not pointers
	copyGrades := &grades
	copyGrades[2] = 100

	fmt.Println(*copyGrades)
	fmt.Println(grades)

	// Arrays are used to back slices
	gradesSlice := []int32{97, 85, 99}
	fmt.Printf("%T, %T\n", grades, gradesSlice)
	fmt.Println(len(gradesSlice), cap(gradesSlice))

	// Slice is a projection of an array
	s1 := grades[0:2]
	s2 := gradesSlice[0:2]
	fmt.Println(s1, s2)

	// Slicing works with both arrays and slices
	// Slices are just a projection of the base array or slice

	a := make([]int, 3) // This also creates a slice
	fmt.Println(a, len(a), cap(a))

	b := []int{}
	b = append(b, 1, 4, 5, 6, 7, 8)
	fmt.Println(b, len(b), cap(b))

	// Concatenate
	b = append(b, []int{9, 10, 11}...)
	fmt.Println(b, len(b), cap(b))

	// Shift operaion
	c := b[1:]
	fmt.Println(c)

	// Remove element from middle??
	d := append(c[0:3], c[4:]...)
	fmt.Println(d)
	fmt.Println(b) // Basically b is fucked because slices are still pointing to the original slice
	// We will need loops to solve this issue

	fmt.Println("-----------------ARRAYS---------------------")
}
