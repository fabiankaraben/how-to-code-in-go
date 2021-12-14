package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Por favor, escribe tu nombre.")
	var name string
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	fmt.Printf("Hola, %s! Yo soy Go!", name)
}
