package chapters

import (
	"fmt"
	"reflect"
)

// MapsAndStructs : Exposed
func MapsAndStructs() {
	fmt.Println("+++++++++++++++++MapsAndStructs+++++++++++++++++++++")
	statePopulation := map[string]int{
		"California": 231232,
		"Texas":      343423432,
		"Florida":    34324,
	}
	fmt.Println(statePopulation)
	// Not the best print output
	// A lot of options with the key types but not infinite
	// Keys should be testable for equality
	// strings, int, bool, string, pointer, interfaces, array, structs, channels are testable
	// slices and maps cannot be used as a key of a map

	// m := map[[]int]string{} This is invalid

	// We can also use the make function to make a map
	// This is for where we don't have the data at the time of declaration
	statePop := make(map[string]int)
	statePop = map[string]int{
		"California": 231232,
		"Texas":      343423432,
		"Florida":    34324,
	}
	fmt.Println(statePop)

	// Like a simple dictionary add a value
	statePop["India"] = 132323
	fmt.Println(statePop["India"])

	// Delete entry from map
	delete(statePop, "India")

	// But this key will still have the initial value of int
	fmt.Println(statePop["India"]) // This will print 0

	// To check if the key is in the map
	pop, ok := statePop["India"]
	fmt.Println("pop, ok : ", pop, ok)

	pop, ok = statePop["Florida"]
	fmt.Println("pop, ok : ", pop, ok)

	// Number of entries
	fmt.Println(len(statePop))

	// Maps are passed as reference which means that values changed in one place is changed
	// everywhere it is copied

	// #############################

	// Structs are used to hold non homogenous data
	// Almost like data classes
	// This is type declaraion
	// Visibility outside the PACKAGE can be controlled at a variable level by making first letter capital
	type Doctor struct {
		number     int
		actorName  string
		Companions []string
	}
	// Order of keys in the declaration of a struct can be in any order
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		Companions: []string{
			"Tera",
			"Mera",
			"Apna",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.Companions[1])

	// What is anonymous struct

	// In some situations you need an anonymous struct
	// This basically means that the type is not defined
	// Only the variable is defined
	// Used where you need to structure your data in a way but don't have a particular type defined
	// Maybe for return statements of functions where you ant to send a subset of data
	// It'll be a pain to define the return types of each function if it changes
	aPatient := struct {
		name string
	}{
		name: "Ramanujan The Great",
	}
	fmt.Println(aPatient)

	// Unline maps - structs are value types, just like int, string...
	// Hence the they are copied by value and you can pass around copies

	fmt.Println(*&aPatient)

	Compositions()
	Tags()

	fmt.Println("-----------------MapsAndStructs---------------------")
}

// Compositions : Exposed
func Compositions() {
	// Through embedding
	// Types can be declared without being used. As they are NOT variables
	fmt.Println("+++++++++++++++++MapsAndStructs/Compositions+++++++++++++++++++++")
	type Animal struct {
		Name   string
		Origin string
	}

	type Bird struct {
		Animal // This is different than BaseAnimal Animal.
		Speed  float32
		CanFly bool
	}

	// The above is embedded into Bird
	// "A bird HAS an animal"
	// It's different than inheritance where "A bird IS an animal"

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.Speed = 48
	b.CanFly = true
	fmt.Printf("%v, %T\n", b, b)
	// In this case we don't have to know the internal structure of the Bird
	// Like we didn't have to see Name and Origin any differently

	// But for literal syntax we have to use the internal structure
	c := Bird{
		Animal: Animal{Name: "Emu", Origin: "Australia"},
		Speed:  25,
		CanFly: false,
	}
	fmt.Println(c)
	fmt.Println("-----------------MapsAndStructs/Compositions---------------------")
}

// Tags exported
func Tags() {

	fmt.Println("+++++++++++++++++MapsAndStructs/Tags+++++++++++++++++++++")

	// Tags are used to describe some specific fields inside a type
	type Animal struct {
		Name   string `required max:"100"`
		Origin string
	}
	// Format
	// Backtiks
	// Space delimited subtags
	// Maybe a key value relationship. keys:"value"

	// Now how do we get at it
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
	fmt.Println(field.Type)
	fmt.Println(t.NumField())

	// What to do with the tag and the field is upto the framework to figure out
	// Framework - like a data validation framework in API definitions

	fmt.Println("-----------------MapsAndStructs/Tags---------------------")

}
