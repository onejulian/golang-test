package testunitario

import "testing"

// func TestSumar(t *testing.T) {
// 	sum := Sumar(5, 5)

// 	if sum != 10 {
// 		t.Errorf("Suma incorrecta, tiene %d y se esperaban %d", sum, 10)
// 	} else {
// 		t.Log("Test finalizado correctamente")
// 	}
// }

func TestSumar(t *testing.T){
	tabla := []struct{
		a int
		b int
		c int
	}{
		{1,2,3},
		{2,2,4},
		{25,25,50},
	}
	for _, item := range tabla{
		sum := Sumar(item.a, item.b)
		if sum != item.c {
			t.Errorf("Suma incorrecta, tiene %d y se esperaba %d", sum, item.c)
		}
	}
}

func TestGetMax(t *testing.T){
	tabla := []struct{
		a int
		b int
		c int
	}{
		{6,2,6},
		{5,9,9},
		{29,26,29},
	}
	for _, item := range tabla{
		max := GetMax(item.a, item.b)
		if max != item.c {
			t.Errorf("Valor incorrecto, resulta %d y se esperaba %d", max, item.c)
		}
	}
}

func TestFibonacci(t *testing.T){
	tabla := []struct{
		a int
		b int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}
	for _, item := range tabla{
		fibo := Fibonacci(item.a)
		if fibo != item.b {
			t.Errorf("Valor incorrecto, resulta %d y se esperaba %d", fibo, item.b)
		}
	}
}