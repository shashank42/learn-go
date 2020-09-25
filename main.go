package main

import "fmt"

// import "../first-go-app/chapters"
// https://github.com/golang/go/wiki/GithubCodeLayout
// import "../first-go-app/chapters"
import "github.com/shashank42/learn-go/chapters"

func main() {
	fmt.Println("Hello Go!")
	chapters.Variables()
	chapters.Primitives()
	chapters.Constants()
	chapters.Arrays()
	chapters.MapsAndStructs()
	chapters.ControlFlow()
	chapters.Looping()
	chapters.DeferAndPanic()
	chapters.Pointers()
	chapters.Functions()
	chapters.Interfaces()
	chapters.Concurrency()
	chapters.Channels()

}
