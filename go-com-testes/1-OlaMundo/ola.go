package main

import "fmt"

// Ola retorna uma saudação no formato "Olá, <nome>!"
func Ola(nome string) string {
	return "Olá, " + nome + "!"
}

func main() {
	fmt.Println(Ola("Daniel"))
}
