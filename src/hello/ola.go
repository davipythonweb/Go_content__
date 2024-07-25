package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Eliote"
	version := 1.1
	age := 46
	fmt.Println("Olá, developer", name, ", sua Idade eh:", age)
	fmt.Println("Este programa está na versão", version)

	fmt.Println("O tipo da variavel name eh:", reflect.TypeOf(name))
	fmt.Println("O tipo da variavel age eh:", reflect.TypeOf(age))
	fmt.Println("O tipo da variavel version eh:", reflect.TypeOf(version))
}

// sem o perador de inferencia de variavel curto

// func main() {
// 	var name = "Eliote"
// 	var version = 1.1
// 	var age = 46
// 	fmt.Println("Olá, developer", name, ", sua Idade eh:", age)
// 	fmt.Println("Este programa está na versão", version)

// 	fmt.Println("O tipo da variavel name eh:", reflect.TypeOf(name))
// 	fmt.Println("O tipo da variavel age eh:", reflect.TypeOf(age))
// 	fmt.Println("O tipo da variavel version eh:", reflect.TypeOf(version))
// }
