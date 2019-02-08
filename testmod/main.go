package main

import (
	"fmt"

	"github.com/nareshkumarthota/rootrepo/testpkg"
)

func main() {
	fmt.Println("i am in main tesmod package")
	testpkg.Display()
}
