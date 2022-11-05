package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	t.Run("particionarLista", func(t *testing.T) {
		lista := []string{"a", "b", "c", "d", "e"}

		resultado := particionarLista(lista, 5)
		esperado := 5

		if len(resultado) == esperado {
			t.Errorf("resultado %v, esperado %d, dado %v", len(resultado), esperado, lista)
		}
	})

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
