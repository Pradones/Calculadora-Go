package main

import "testing"

func TestSoma(t *testing.T) {
	teste := soma(3, 5, 7)
	resultado := 15

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestSoma2(t *testing.T) {
	teste := soma(3, 5, 7)
	resultado := 18

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TesteMult(t *testing.T) {
	teste := multiplic(10, 20)
	resultado := 200
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor obtido:", teste)
	}
}

func TesteMult2(t *testing.T) {
	teste := multiplic(10, 20)
	resultado := 300
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor obtido:", teste)
	}
}

func TesteSubtrai(t *testing.T) {
	teste := subtrai(5, 15)
	resultado := -10
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor obtido:", teste)
	}
}

func TesteSubtrai2(t *testing.T) {
	teste := subtrai(5, 15)
	resultado := -20
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor obtido:", teste)
	}
}

func TesteDivide(t *testing.T) {
	teste := divide(5, 10)
	resultado := 2
	if teste != resultado {
		t.Error("Valor esperado", resultado, "Valor obtido:", teste)
	}
}

func TesteDivide2(t *testing.T) {
	teste := divide(5, 10)
	resultado := 4
	if teste != resultado {
		t.Error("Valor esperado", resultado, "Valor obtido:", teste)
	}
}
