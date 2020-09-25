package chapters

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter int = 0
var m = sync.RWMutex{} // Infinite number of readers but only one writer

// Concurrency : Introduce concurrency
func Concurrency() {
	runtime.GOMAXPROCS(4)
	fmt.Println("+++++++++++++++++Concurrency+++++++++++++++++++++")
	wg.Add(1)
	go sayHello()
	// Spin off a green thread instead of a OS thread
	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHelloCounter()
		m.Lock()
		go increment()
	}
	wg.Wait()

	// fmt.Println(wg)

	// Don't create goroutinues in libraries. Let consumer control concurrency
	// GoRoutinues may run forever

	fmt.Println("-----------------Concurrency---------------------")
}

func sayHello() {
	fmt.Println("Hello")
	wg.Done()
}

func sayHelloCounter() {

	fmt.Println("Hello :", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {

	counter++
	m.Unlock()
	wg.Done()
}
