package main 

import "fmt"

func main() {
	Tabla := tabla(2)
	for i:= 1; i<= 10; i++ {
		fmt.Printf("2 x %v = %v \n",i,Tabla())
	}
}

// Clousure
func tabla(valor int) func() int {
	numero := valor
	secuencia:= 0
	return func() int {
		secuencia ++
		return numero*secuencia
	}
}