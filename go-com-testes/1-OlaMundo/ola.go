package main

import "fmt"

const espanhol = "espanhol"
const frances = "francês"
const prefixoOlaPortuges = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

// Ola retorna uma saudação no formato "Olá, <nome>!" dependendo da linguagem
func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return prefixoSaudacao(idioma) + nome
}

/*
 *	Função privada... funções privadas começam com letras minusculas
 *  enquanto publicas começam com letras maiusculas
 */
func prefixoSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case frances:
		prefixo = prefixoOlaFrances
	default:
		prefixo = prefixoOlaPortuges
	}
	return
}

func main() {
	fmt.Println(Ola("Daniel", ""))
}

// Ciclo TDD:
// Escreva um teste que falhe e veja-o falhar, para que saibamos que escrevemos um teste relevante para nossos requisitos e vimos que ele produz uma descrição da falha fácil de entender
// Escrever a menor quantidade de código para fazer o teste passar, para que saibamos que temos software funcionando
// Em seguida, refatorar, tendo a segurança de nossos testes para garantir que tenhamos um código bem feito e fácil de trabalhar
