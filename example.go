/*
	Un CRUD completo de GoLang y MySQL
	@author parzibyte
*/
package main

import (
	"bufio"                            // Leer líneas incluso si tienen espacios
	"database/sql"                     // Interactuar con bases de datos
	"fmt"                              // Imprimir mensajes y esas cosas
	_ "github.com/go-sql-driver/mysql" // La librería que nos permite conectar a MySQL
	"os"                               // El búfer, para leer desde la terminal con os.Stdin
)

type Contacto struct {
	Nombre, Direccion, CorreoElectronico string
	Id                                   int
}

func obtenerBaseDeDatos() (db *sql.DB, e error) {
	usuario := "root"
	pass := ""
	host := "tcp(127.0.0.1:3306)"
	nombreBaseDeDatos := "prueba"
	// Debe tener la forma usuario:contraseña@protocolo(host:puerto)/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	creditos := `==========================================================
	                CRUD de MySQL y GO
                                           
                                __ __          __         
	.-----.---.-.----.-----|__|  |--.--.--|  |_.-----.
	|  _  |  _  |   _|-- __|  |  _  |  |  |   _|  -__|
	|   __|___._|__| |_____|__|_____|___  |____|_____|
	|__|                            |_____|           
==========================================================`
	fmt.Println(creditos)
	menu := `¿Qué deseas hacer?
[1] -- Insertar
[2] -- Mostrar
[3] -- Actualizar
[4] -- Eliminar
[5] -- Salir
----->	`
	var eleccion int
	var c Contacto
	for eleccion != 5 {
		fmt.Print(menu)
		fmt.Scanln(&eleccion)
		scanner := bufio.NewScanner(os.Stdin)
		switch eleccion {
		case 1:
			fmt.Println("Ingresa el nombre:")
			if scanner.Scan() {
				c.Nombre = scanner.Text()
			}
			fmt.Println("Ingresa la dirección:")
			if scanner.Scan() {
				c.Direccion = scanner.Text()
			}
			fmt.Println("Ingresa el correo electrónico:")
			if scanner.Scan() {
				c.CorreoElectronico = scanner.Text()
			}
			err := insertar(c)
			if err != nil {
				fmt.Printf("Error insertando: %v", err)
			} else {
				fmt.Println("Insertado correctamente")
			}
		case 2:
			contactos, err := obtenerContactos()
			if err != nil {
				fmt.Printf("Error obteniendo contactos: %v", err)
			} else {
				for _, contacto := range contactos {
					fmt.Println("====================")
					fmt.Printf("Id: %d\n", contacto.Id)
					fmt.Printf("Nombre: %s\n", contacto.Nombre)
					fmt.Printf("Dirección: %s\n", contacto.Direccion)
					fmt.Printf("E-mail: %s\n", contacto.CorreoElectronico)
				}
			}
		case 3:
			fmt.Println("Ingresa el id:")
			fmt.Scanln(&c.Id)
			fmt.Println("Ingresa el nuevo nombre:")
			if scanner.Scan() {
				c.Nombre = scanner.Text()
			}
			fmt.Println("Ingresa la nueva dirección:")
			if scanner.Scan() {
				c.Direccion = scanner.Text()
			}
			fmt.Println("Ingresa el nuevo correo electrónico:")
			if scanner.Scan() {
				c.CorreoElectronico = scanner.Text()
			}
			err := actualizar(c)
			if err != nil {
				fmt.Printf("Error actualizando: %v", err)
			} else {
				fmt.Println("Actualizado correctamente")
			}
		case 4:
			fmt.Println("Ingresa el ID del contacto que deseas eliminar:")
			fmt.Scanln(&c.Id)
			err := eliminar(c)
			if err != nil {
				fmt.Printf("Error eliminando: %v", err)
			} else {
				fmt.Println("Eliminado correctamente")
			}
		}
	}
}

func eliminar(c Contacto) error {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("DELETE FROM agenda WHERE id = ?")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()

	_, err = sentenciaPreparada.Exec(c.Id)
	if err != nil {
		return err
	}
	return nil
}

func insertar(c Contacto) (e error) {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	// Preparamos para prevenir inyecciones SQL
	sentenciaPreparada, err := db.Prepare("INSERT INTO agenda (nombre, direccion, correo_electronico) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(c.Nombre, c.Direccion, c.CorreoElectronico)
	if err != nil {
		return err
	}
	return nil
}

func obtenerContactos() ([]Contacto, error) {
	contactos := []Contacto{}
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT id, nombre, direccion, correo_electronico FROM agenda")

	if err != nil {
		return nil, err
	}
	// Si llegamos aquí, significa que no ocurrió ningún error
	defer filas.Close()

	// Aquí vamos a "mapear" lo que traiga la consulta en el while de más abajo
	var c Contacto

	// Recorrer todas las filas, en un "while"
	for filas.Next() {
		err = filas.Scan(&c.Id, &c.Nombre, &c.Direccion, &c.CorreoElectronico)
		// Al escanear puede haber un error
		if err != nil {
			return nil, err
		}
		// Y si no, entonces agregamos lo leído al arreglo
		contactos = append(contactos, c)
	}
	// Vacío o no, regresamos el arreglo de contactos
	return contactos, nil
}

func actualizar(c Contacto) error {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("UPDATE agenda SET nombre = ?, direccion = ?, correo_electronico = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Pasar argumentos en el mismo orden que la consulta
	_, err = sentenciaPreparada.Exec(c.Nombre, c.Direccion, c.CorreoElectronico, c.Id)
	return err // Ya sea nil o sea un error, lo manejaremos desde donde hacemos la llamada
}