package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 2
const delay = 5

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
		case 0:
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
	fmt.Println("")

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := []string{"https://solyd.com.br/",
		"https://olhardigital.com.br/", "https://pato.academy/"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ": ", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")

	}

	fmt.Println("")

}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:",
			resp.StatusCode)
	}
}
