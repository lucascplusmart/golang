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

func lerArquivo(filename string) (data []byte) {
	// metodo para ler o arquivo
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panicf("Falha ao ler o arquivo: %s", err)
	}

	fmt.Printf("\nNome do arquivo: %s", filename)
	fmt.Printf("\nTamanho: %d bytes\n", len(data))

	return data
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

	var (
		controleRotina sync.WaitGroup
		valorBusca     string
	)

	data := lerArquivo("teste.txt")
	lista := obterLista(string(data))
	superLista := particionarLista(lista, MAX_GOROTINAS)
	canal := make(chan string, 0)

	fmt.Println("Informe o valor de busca:")
	n, erro := fmt.Scanln(&valorBusca)
	if erro != nil {
		log.Println("Erro ao ler valor de busca:", erro)
		return
	}

	for i := range superLista {
		controleRotina.Add(1)
		go processarBuscar(superLista[i], &controleRotina, valorBusca, canal)
	}

	go func() {
		controleRotina.Wait()
		close(canal)
	}()

	valor, existeValor := <-canal

	if existeValor {
		fmt.Println("valor encontrado:", valor, n)
	} else {
		fmt.Println("valor não encontrado!")
	}

}
