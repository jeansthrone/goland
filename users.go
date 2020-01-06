package main

import (
	"fmt"
	"prueba/conection"
)

func main() {
	// Insertar datos

	c := conection.User{
		Name:     "Mia",
		Lastname: "fol",
	}

	err := conection.Insertar(c)

	if err != nil {
		fmt.Printf("Error insertando: %v", err)
	} else {
		fmt.Println("\n Insertado correctamente\n")
	}

	// Actualizar Datos

	newuser := conection.User{
		Name:     "Mihsaha",
		Lastname: "Trusted",
		Id:       52,
	}
	acterr := conection.Actualizar(newuser)
	if acterr != nil {
		fmt.Printf("Error actualizando: %v", acterr)
	} else {
		fmt.Println("Actualizado correctamente")
	}

	// Eliminar Datos

	EliminarUser := conection.User{
		Id: 52,
	}
	err := conection.Eliminar(EliminarUser)
	if err != nil {
		fmt.Printf("Error al eliminar: %v", err)
	} else {
		fmt.Println("Eliminado correctamente")
	}

	// Mostrar Datos

	users, err := conection.GetUsuarios()
	if err != nil {
		fmt.Printf("Error obteniendo usuarios: %v", err)
		return
	}
	for _, user := range users {
		fmt.Printf("Nombre: %s\n", user.Name+" "+user.Lastname)
	}

}
