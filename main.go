package main

import (
	"bufio" // Leer líneas incluso si tienen espacios
	"fmt"
	"os" // El búfer, para leer desde la terminal con os.Stdin
	"prueba/conection"
)

func main() {

	menu := `¿Qué deseas hacer?
[1] -- Insertar
[2] -- Mostrar
[3] -- Actualizar
[4] -- Eliminar
[5] -- Salir
----->	`
	var eleccion int
	var c conection.User
	for eleccion != 5 {
		fmt.Print(menu)
		fmt.Scanln(&eleccion)
		scanner := bufio.NewScanner(os.Stdin)
		switch eleccion {
		case 1:
			fmt.Println("Ingresa el nombre:")
			if scanner.Scan() {
				c.Name = scanner.Text()
			}
			fmt.Println("Ingresa el apellido:")
			if scanner.Scan() {
				c.Lastname = scanner.Text()
			}
			err := conection.Insertar(c)

			if err != nil {
				fmt.Printf("\nError insertando: %v\n", err)
			} else {
				fmt.Println("\n Insertado correctamente\n")
			}
		case 2:
			users, err := conection.GetUsuarios()
			if err != nil {
				fmt.Printf("\nError obteniendo usuarios: %v\n", err)
				return
			} else {
				fmt.Println("\n")
				for _, user := range users {
					fmt.Println("====================")
					fmt.Printf("Nombre: %s\n", user.Name+" "+user.Lastname)
				}
				fmt.Println("====================")
				fmt.Println("\n")
			}
		case 3:
			fmt.Println("\nIngresa el id:")
			fmt.Scanln(&c.Id)
			fmt.Println("Ingresa el nuevo nombre:")
			if scanner.Scan() {
				c.Name = scanner.Text()
			}
			fmt.Println("Ingresa el nuevo apellido:")
			if scanner.Scan() {
				c.Lastname = scanner.Text()
			}

			err := conection.Actualizar(c)
			if err != nil {
				fmt.Printf("\nError actualizando: %v\n", err)
			} else {
				fmt.Println("\nActualizado correctamente\n")
			}
		case 4:
			fmt.Println("\nIngresa el ID del usuario que deseas eliminar:")
			fmt.Scanln(&c.Id)
			err := conection.Eliminar(c)
			if err != nil {
				fmt.Printf("\nError eliminando: %v\n", err)
			} else {
				fmt.Println("\nEliminado correctamente\n")
			}
		}
	}
}
