package main

import "fmt"

func main() {

	//Clousure

	// Tabla := tabla(2)
	// for i:= 1; i<= 10; i++ {
	// 	fmt.Printf("2 x %v = %v \n",i,Tabla())
	// }

	// Estructura anidada
	categoria := Categoria{Id: 1, Nombre: "Categoria 1", Slug: "categoria-1"}
	producto := Producto{Id: 1, Nombre: "Mesa de ordenador", Slug:"mesa-de-ordenador",Precio:1234,CategoriaId: categoria}
	fmt.Printf("%+v \n",producto)
}

// Clousure
// func tabla(valor int) func() int {
// 	numero := valor
// 	secuencia := 0
// 	return func() int {
// 		secuencia++
// 		return numero * secuencia
// 	}
// }

type Categoria struct {
	Id     int
	Nombre string
	Slug   string
}
type Producto struct {
	Id          int
	Nombre      string
	Slug        string
	Precio      int
	CategoriaId Categoria
}
