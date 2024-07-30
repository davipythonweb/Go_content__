package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 3:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Ops! Comando desconhecido.")
			os.Exit(-1) //o numero -1 indica que ocorreu algum erro no programa.
		}
	}
}

func exibeIntroducao() {
	name := "Eliote"
	version := 1.1
	fmt.Println("Olá, developer", name)
	fmt.Println("Este programa está na versão", version)
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br/"
	resp, _ := http.Get(site)
	// fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:",
			resp.StatusCode)
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
