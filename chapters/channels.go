package chapters

import (
	"fmt"
	"sync"
	"time"
)

// Channels : Exposed
func Channels() {
	// Pass data between goroutinues
	fmt.Println("+++++++++++++++++Channels+++++++++++++++++++++")
	// Basics
	basicChannels()

	// Restricting data flows
	restrictedChannels()

	// Buffered channels
	// Closing channels
	// For range with channels
	bufferedClosingChannels()

	// Select statements for channels
	logging()

	fmt.Println("-----------------Channels---------------------")
}

var wg1 = sync.WaitGroup{}

func basicChannels() {
	// By default there is only one element of data in channel
	ch := make(chan int)
	wg1.Add(2)

	go func() {
		i := <-ch
		fmt.Println(i)
		wg1.Done()
	}()
	go func() {
		i := 42
		ch <- i // Passed by value
		i = 27
		wg1.Done()
	}()
	wg1.Wait()

}

func restrictedChannels() {
	fmt.Println("+++++++++++++++++Channels/restrictedChannels+++++++++++++++++++++")
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) { // There is some casing going on as you are passing bidirectional channel
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()
	fmt.Println("-----------------Channels/restrictedChannels---------------------")
}

func bufferedClosingChannels() {
	fmt.Println("+++++++++++++++++Channels/bufferedChannels+++++++++++++++++++++")
	ch := make(chan int, 50)
	wg1.Add(2)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		i = <-ch
		fmt.Println(i)
		for j := range ch {
			fmt.Println(j)
		}
		for {
			if k, ok := <-ch; ok {
				fmt.Println(k)
			} else {
				break
			}
		}
		wg1.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 54
		for j := 0; j < 10; j++ {
			ch <- j
		}
		close(ch) // No way to detect really // Only way to recover from panic
		wg1.Done()
	}(ch)
	wg1.Wait()
	fmt.Println("-----------------Channels/bufferedChannels---------------------")
}

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // Purely used for message signaling

// Signal only channel. 0 memory allocation
// Can't send any data through

func logging() {
	go logger()

	// This is sensible
	// defer func() {
	// 	close(logCh)
	// }()

	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	doneCh <- struct{}{}
}

func logger() {
	// This can be used
	//for entry := range logCh {
	//	fmt.Println(entry)
	//}

	for {
		select {
		case entry := <-logCh:
			fmt.Println(entry)
		case <-doneCh:
			fmt.Println("Closing logger")
			break
		}
	}
}
