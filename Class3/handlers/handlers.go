package handlers

import (
	"Class_2_mysql_driver/conectar"
	"Class_2_mysql_driver/modelos"
	"fmt"
)

func Listar() {
	conectar.Conectar()
	sql := "select id,nombre,correo,telefono from clientes order by id desc;"
	datos, err := conectar.Db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}

	defer conectar.CerrarConexion()
	clientes := modelos.Clientes{}
	for datos.Next() {
		dato := modelos.Cliente{}
		datos.Scan(&dato.Id,&dato.Nombre,&dato.Correo,&dato.Telefono)
		clientes = append(clientes,dato)
	}
	fmt.Println(clientes)
}

func ListarPorId(id int) {
	conectar.Conectar()
	sql := "select id,nombre,correo,telefono from clientes where id=?;"
	datos, err := conectar.Db.Query(sql,id)
	if err != nil {
		fmt.Println(err)
	}

	defer conectar.CerrarConexion()
	for datos.Next() {
		var dato modelos.Cliente
		datos.Scan(&dato.Id,&dato.Nombre,&dato.Correo,&dato.Telefono)
		fmt.Println(dato)
	}
}