package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Daniel")
	esperado := "Olá, Daniel!"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
