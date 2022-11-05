package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
	"sync"
)

const MAX_GOROTINAS uint = 100 // maximo de go rotinas que podem ser criadas para processar os dados no sistema

func particionarLista(s []string, sub uint) [][]string {
	size := int(math.Ceil(float64(len(s)) / float64(sub)))
	var list = make([][]string, 0)
	var j int

	for i := 0; i < len(s); i += size {
		j += size
		if j > len(s) {
			j = len(s)
		}

		list = append(list, s[i:j])
	}
	return list
}

func lerArquivo(caminho, nomeArquivo *string) (erro error, data []byte) {
	var arquivo = *nomeArquivo
	if caminho != nil || *caminho != "" {
		arquivo = *caminho + *nomeArquivo
	}

	// metodo para ler o arquivo
	data, erro = ioutil.ReadFile(arquivo)
	if erro != nil {
		return erro, nil
	}

	fmt.Printf("\nNome do arquivo: %s", *nomeArquivo)
	fmt.Printf("\nTamanho: %d bytes\n", len(data))

	return nil, data
}

func obterLista(s string) []string {
	return strings.Split(s, ",")
}

func processarBuscar(lista []string, controleRotina *sync.WaitGroup, valorBusca string, canal chan string) {
	defer controleRotina.Done()
	for i := range lista {
		if lista[i] == valorBusca {
			canal <- lista[i]
		}
	}

}

func main() {
	var valorBusca, caminhoArquivo, nomeArquivo string

	fmt.Println("Informe o caminho do arquivo:")
	_, erro := fmt.Scanln(&caminhoArquivo)
	if erro != nil {
		log.Println("Erro ao ler o caminho do arquivo:", erro)
		return
	}

	fmt.Println("Informe o nome do arquivo:")
	_, erro = fmt.Scanln(&nomeArquivo)
	if erro != nil {
		log.Println("Erro ao ler o nome do arquivo:", erro)
		return
	}

	// manipulando os dados do arquivo
	erro, data := lerArquivo(&caminhoArquivo, &nomeArquivo)
	if erro != nil {
		log.Panic("Erro ao ler o nome do arquivo:", erro)
	}
	lista := obterLista(string(data))

	// Capturando os valores de busca
	fmt.Println("Informe o valor de busca:")
	_, erro = fmt.Scanln(&valorBusca)
	if erro != nil {
		log.Println("Erro ao ler valor de busca:", erro)
		return
	}

	var (
		valor       string
		existeValor bool
	)

	// buscando os dados
	for i := range lista {
		if lista[i] == valorBusca {
			valor = valorBusca
			existeValor = true
		}
	}

	if existeValor {
		fmt.Println("valor encontrado:", valor)
	} else {
		fmt.Println("valor não encontrado!")
	}
}
