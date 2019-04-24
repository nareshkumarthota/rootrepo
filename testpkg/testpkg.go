package testpkg

import "fmt"

func init() {
	Display()
}

// Display display method
func Display() {
	fmt.Println("test package repo display method()")
	fmt.Println("new method added")
}
