package main

import (
	"fmt"
)

func main() {
	name := "Eliote"
	version := 1.1
	fmt.Println("Olá, developer", name)
	fmt.Println("Este programa está na versão", version)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Sair do Programa")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi", comando)

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 3:
		fmt.Println("Saindo do Programa...")
	default:
		fmt.Println("Ops! Comando desconhecido.")
	}
}

// func main() {
// 	name := "Eliote"
// 	version := 1.1
// 	fmt.Println("Olá, developer", name)
// 	fmt.Println("Este programa está na versão", version)

// 	fmt.Println("1- Iniciar Monitoramento")
// 	fmt.Println("2- Exibir Logs")
// 	fmt.Println("3- Sair do Programa")

// 	var comando int
// 	// fmt.Scanf("%d", &comando)
// 	//ou
// 	fmt.Scan(&comando) // o -> & , serve para pegar o endereço ou ponteiro da variavel.
// 	fmt.Println("O endereço ou ponteiro da variavel comando eh:", &comando)

// 	fmt.Println("O comando escolhido foi", comando)

// 	// if comando == 1 {
// 	// 	fmt.Println("Monitorando...")
// 	// } else if comando == 2 {
// 	// 	fmt.Println("Exibindo Logs...")

// 	// } else if comando == 3 {
// 	// 	fmt.Println("Saindo do Programa...")
// 	// } else {
// 	// 	fmt.Println("Ops! Comando desconhecido.")
// 	// }

// 	// usando switch
// 	switch comando {
// 	case 1:
// 		fmt.Println("Monitorando...")
// 	case 2:
// 		fmt.Println("Exibindo Logs...")
// 	case 3:
// 		fmt.Println("Saindo do Programa...")
// 	default:
// 		fmt.Println("Ops! Comando desconhecido.")
// 	}
// }

// func main() {
// 	name := "Eliote"
// 	version := 1.1
// 	age := 46
// 	fmt.Println("Olá, developer", name, ", sua Idade eh:", age)
// 	fmt.Println("Este programa está na versão", version)

// 	fmt.Println("O tipo da variavel name eh:", reflect.TypeOf(name))
// 	fmt.Println("O tipo da variavel age eh:", reflect.TypeOf(age))
// 	fmt.Println("O tipo da variavel version eh:", reflect.TypeOf(version))
// }

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
