package chapters

import (
	"bytes"
	"fmt"
	"io"
)

// Interfaces : Exported
func Interfaces() {

	fmt.Println("+++++++++++++++++Interfaces+++++++++++++++++++++")
	// Interfaces describe behavior
	// Structs describe data
	// Basics
	basic()

	// Compose interfaces
	complicated() // Need to implement all of the functions in the embedded interfaces

	// Type conversion
	typeConversion()
	// 		Empty interface
	emptyInterface()
	//		Type switches
	typeSwitch()

	// Implement interface
	//		By value type
	//		By pointers
	// Types that are assigned to the interface as pointers can be implemented by both
	// value recievers as well as pointer recievers
	// While Types that are assigned to the interface as values can be implemented by
	// only value recievers
	// Type assignment eg. var obj WriterCloser = &BufferedWriterCloser
	// Value reciept eg. func (bwc *BufferedWriterCloser) Close() error

	// Best practices
	// 1. Use many, small interfaces instead of one big monolithic
	//		eg. io.Writer, io.Reader, interface{}
	// 2. Don't export interfaces for types that will be consumed
	// 3. Do export interfaces for types that will be used by packages
	// 4. Design functions and methods to recieve interfaces whenever possible

	fmt.Println("-----------------Interfaces---------------------")
}

// Writer : Exposed
type Writer interface {
	Write([]byte) (int, error)
}

// ConsoleWriter exposed
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func basic() {

	// Implicit implementation
	var w Writer = ConsoleWriter{} // This interface initialization can be anything
	w.Write([]byte("Hello guys!"))
	// Any type can implement an interface. Here a struct is implementing the interface because the Write function is part
	// of the function that is defined using the struct ConsoleWriter

}

// Closer to create the interface that captures teh Closing behaviour
type Closer interface {
	Close() error
}

// WriterCloser to create the interface that composes of Writer and Closer
type WriterCloser interface {
	Writer
	Closer
}

// BufferedWriterCloser to implement the WriterCloser interface
type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

// Write implemented for BufferedWriterCloser as an implementation of WriterCloser interface
func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 0 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v), " : ")
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

// Close implemented for BufferedWriterCloser as an implementation of WriterCloser interface
func (bwc *BufferedWriterCloser) Close() error {
	fmt.Println("Close the buffer")
	for bwc.buffer.Len() > 1 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

// NewBufferedWriterCloser : Constructor function
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

func complicated() {

	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello shashank, how is your day"))
	wc.Close()
}

func typeConversion() {
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello type conversion"))
	wc.Close()

	// This is the type conversion step
	// After this you will have access to the internal buffer of the BufferedWriterCloser
	// When it was only a WriterCloser, Go did not have access to the internal data of the
	// struct
	bwc, _ := wc.(*BufferedWriterCloser)
	fmt.Printf("%v, %T\n", bwc, bwc)

	// How about conversion to something that is not allowed
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed", ok)
	}
}

func emptyInterface() {
	// Very often a type switch will be paired with an empty interface
	var myObj interface{} = NewBufferedWriterCloser()
	if wc, ok := myObj.(WriterCloser); ok {
		wc.Write([]byte("This is using empty interface"))
		wc.Close()
	}
	r, ok := myObj.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}
}

func typeSwitch() {
	var i interface{} = NewBufferedWriterCloser()
	switch i.(type) {
	case int:
		fmt.Println("i is integer")
	case string:
		fmt.Println("i is string")
	case *BufferedWriterCloser:
		fmt.Println("i is BufferedWriterCloser")
	default:
		fmt.Println("Who knows what is what")
	}
}
