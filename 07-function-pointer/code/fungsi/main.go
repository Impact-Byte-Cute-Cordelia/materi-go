package main

import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"Foo", "Bar"}

	fmt.Println(strings.Join(names, " "))

	paggilFungsiLain()
}

func paggilFungsiLain() {
	fmt.Println("Hello")
}
