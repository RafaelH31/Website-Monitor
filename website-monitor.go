package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimirLogs()
		case 0:
			fmt.Println("Saindo do programa")

			fmt.Println(comando)

			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	fmt.Println("Bem vindo ao Monitorador de Sites")
	fmt.Println("Este programa está na versão 1.22")
}

func exibeMenu() {
	fmt.Println("\n1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scanf("%d\n", &comandoLido)
	fmt.Println("\nO comando escolhido foi", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("\nMonitorando...")

	var urls []string

	obterUrls(&urls)

	fmt.Println("\nURLs a serem monitoradas:")
	for _, site := range urls {
		fmt.Println(site)
	}

	fmt.Print("\nDigite em segundos o intervalo de tempo a cada requisição:")

	var tempoRequisicao int
	fmt.Scanln(&tempoRequisicao)

	fmt.Print("\nDigite em segundos o tempo de monitoramento dos sites:")

	var tempoMoni int
	fmt.Scanln(&tempoMoni)

	for i := 0; i < tempoMoni; i += tempoRequisicao {
		for numero, url := range urls {
			fmt.Println("Testando site", numero, ":", url)
			testaUrl(url)
		}
		time.Sleep(time.Duration(tempoRequisicao) * time.Second)
	}
	fmt.Println("")
}

func obterUrls(urls *[]string) {

	fmt.Print("\nDigite a URL que deseja monitorar (ou 'sair' para encerrar): ")

	for {
		var url string
		fmt.Scanln(&url)

		if url == "sair" {
			break
		}

		if url != "" {
			*urls = append(*urls, url)
		}
	}
}

func testaUrl(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Erro ao realizar a requisição:", err)
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", url, "foi carregado com sucesso!")
		registrarLog(url, true)
	} else {
		fmt.Println("Site:", url, "está com problemas. Status Code:", resp.StatusCode)
		registrarLog(url, false)
	}
}

func registrarLog(url string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + url + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimirLogs() {

	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}
