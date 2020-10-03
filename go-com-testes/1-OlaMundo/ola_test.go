package main

import "testing"

func TestOla(t *testing.T) {
	verificarMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("Diz Olá para as pessoas, dado o idioma. Teste do idioma Espanhol", func(t *testing.T) {
		resultado := Ola("Daniel", "espanhol")
		esperado := "Hola, Daniel"
		verificarMensagemCorreta(t, resultado, esperado)

	})

	t.Run("Diz Olá para as pessoas, dado o idioma. Teste do idioma Francês", func(t *testing.T) {
		resultado := Ola("Daniel", "francês")
		esperado := "Bonjour, Daniel"
		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("Diz 'Olá, mundo' quando uma string vazia for passada nos campos", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, Mundo"
		verificarMensagemCorreta(t, resultado, esperado)
	})
}
