package chapters

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// DeferAndPanic : Exposed
func DeferAndPanic() {
	fmt.Println("+++++++++++++++++DeferAndPanic+++++++++++++++++++++")

	deferFunc()
	practicalDefer()
	weirdDefer()

	//panicFunc()
	//deferAndPanic()
	//deferPanicRecover()
	deeperDeferPanicRecover()

	fmt.Println("-----------------DeferAndPanic---------------------")
}

func deferFunc() {

	fmt.Println("+++++++++++++++++deferFunc+++++++++++++++++++++")
	fmt.Println("start")
	// Any function passed to defer gets run at the end of the function and before return
	defer fmt.Println("middle")
	defer fmt.Println("middle - 2")
	defer fmt.Println("middle - 3")
	// defer functions are called in LIFO
	// Because we use it for closing out resources and makes sense to do it LIFO
	fmt.Println("end")
	fmt.Println("-----------------deferFunc---------------------")

}

func practicalDefer() {
	fmt.Println("+++++++++++++++++practicalDefer+++++++++++++++++++++")
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s, %T", robots, robots)
	fmt.Println("-----------------practicalDefer---------------------")

}

func weirdDefer() {
	a := "start"
	defer fmt.Println(a) // Go locks in the parameter values with the deferred function call
	a = "end"
}

func panicFunc() {
	a, b := 1, 0
	if b == 0 {
		panic("b is 0")
	}
	ans := a / b
	fmt.Println(ans)
}

func deferAndPanic() {
	// First execute the function
	// Then execute any deferred functions
	// Then handle any panic functions
	// Then handle the return values
	// Thus any deferred call to close out resources will be executed
	fmt.Println("Start")
	defer fmt.Println("This was deferred")
	panic("Something bad happened") // Panic happens after defer happens
	fmt.Println("end")

}

func deferPanicRecover() {
	// First execute the function
	// Then execute any deferred functions
	// Then handle any panic functions
	// Then handle the return values
	// Thus any deferred call to close out resources will be executed
	fmt.Println("Start")
	defer func() {
		if err := recover(); err != nil { // Recover function returns nil or the error causing teh application to panic
			log.Println("Error:", err)
		}
	}()
	panic("Something bad happened") // Panic happens after defer happens
	fmt.Println("end")

}

func deeperDeferPanicRecover() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			// panic(err)
		}
	}()
	fmt.Println(recover())
	panic("something bad happened")
	fmt.Println("done panicking") // This will not print.
	// Panicker function stops it right there and executes the defered functions and then panics
	// Functions higher up the call stack can continue based on your condition
	// You can panic inside the recover function
	// It will pass the panic top the Go Runtime which doesn't have built in error handler
}
