package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	t.Run("lerArquivo", func(t *testing.T) {
		nome := "teste.txt"

		erro, _ := lerArquivo(nil, &nome)
		if erro != nil {
			t.Errorf("erro ao ler arquivo")
		}
	})

	t.Run("obterLista", func(t *testing.T) {
		teste := "a,b,c"
		res := obterLista(teste)
		if len(res) != 3 {
			t.Errorf("erro ao ler arquivo")
		}

	})
}
