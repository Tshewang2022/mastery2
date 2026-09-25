package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Hello, World!")
}

func cleanInput(text string) []string {
	return strings.Fields(text)
}
