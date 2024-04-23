package main

import (
	"testing"
)

func TestSoma(t *testing.T) {

	teste := soma(5, 6)
	resultado := 11

	if teste != resultado { 
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}

}

func TestSubtracao(t *testing.T) {

	teste := subtracao(5, 25)
	resultado := 20

	if teste != resultado { 
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}

}

func TestMultiplicacao(t *testing.T) {

	teste := multiplicacao(10, 10)
	resultado := 100

	if teste != resultado { 
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}

}

func TestDivisao(t *testing.T) {

	teste := divisao(2, 40)
	resultado := 20.0

	if teste != resultado { 
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}

}