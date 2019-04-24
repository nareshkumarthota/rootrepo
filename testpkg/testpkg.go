package testpkg

import (
	"fmt"
	"os"
)

func init() {
	Display()
}

// Display display method
func Display() {
	fmt.Println("test package repo display method()")
	fmt.Println("new method added")
}

// NewMethod new functionality implementing
func NewMethod() {
	fmt.Println("hello in new method")
	os.Getwd()
}

// NewMethod1 new functionality implementing
func NewMethod1() {
	fmt.Println("hello in new method1")
	os.Getwd()
}
