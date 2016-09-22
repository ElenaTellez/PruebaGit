package main

import "fmt"

// fib devuelve una funcion sucesiva de los numeros de Fibonacci

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	
	f := fib()
	// La funcion llamada es evaluada de izquierda a derecha
	fmt.Println(f(), f(), f(), f(), f())
}
