package conection

import (
	"database/sql"                     // Interactuar con bases de datos
	"fmt"                              // Imprimir mensajes y esas cosas
	_ "github.com/go-sql-driver/mysql" // La librería que nos permite conectar a MySQL
)

func obtenerBaseDeDatos() (db *sql.DB, e error) {
	usuario := "root"
	pass := ""
	host := "tcp(127.0.0.1:3306)"
	nombreBaseDeDatos := "prueba"
	// Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}

/*func main() {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		fmt.Printf("Error obteniendo base de datos: %v", err)
		return
	}
	// Terminar conexión al terminar función
	defer db.Close()

	// Ahora vemos si tenemos conexión
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error conectando: %v", err)
		return
	}
	// Listo, aquí ya podemos usar a db!
	fmt.Printf("Conectado correctamente")
}*/

type User struct {
	Name, Lastname string
	Id             int
}

//------------------------ Mostrar Datos ------------------------

func GetUsuarios() ([]User, error) {
	users := []User{}
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	filas, err := db.Query("SELECT name, lastname FROM users")
	if err != nil {
		return nil, err
	}

	//egamosaqunificque no ocurió nngún er

	defer filas.Close()

	//Aquí vamos a "mapear" o uetraia la consute e hie de más abajo

	var c User

	//ecorrer todas lls, en u "while"

	for filas.Next() {
		err2 := filas.Scan(&c.Name, &c.Lastname)
		// Al escaneapued aber u error
		if err2 != nil {
			return nil, err2
		}
		// Y si no, entoces agramos lo leído al arreglo

		users = append(users, c)

	}
	// Vacío o no, reesamos el arreglo de contactos

	//Esto no funcionó
	/*var usuario []sting
		for filas.Next() {
			var name, lastname string
			err3 := filas.Scn(&name, &lastname)
			if err3 != nil {
				mt.Println("error3")
		}

			fmt.Println(name + " " + lastname)
		usuario = append(name + " " + lastname)

	}*/

	return users, nil

}

//------------------------ Insertar Datos ------------------------

func Insertar(c User) (e error) {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	// Preparamos para prevenir inyecciones SQL
	sentenciaPreparada, err := db.Prepare("INSERT INTO users(name, lastname) VALUES( ?, ? )")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(c.Name, c.Lastname)
	if err != nil {
		return err
	}
	return nil
}

//------------------------ Actualizar Datos ------------------------

func Actualizar(c User) (e error) {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	// Preparamos para prevenir inyecciones SQL
	sentenciaPreparada, err := db.Prepare("UPDATE users SET name = ?, lastname = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()

	// Pasar argumentos en el mismo orden que la consulta
	_, err = sentenciaPreparada.Exec(c.Name, c.Lastname, c.Id)
	return err // Ya sea nil o sea un error, lo manejaremos desde donde hacemos la llamada

}

//------------------------ Eliminar Datos ------------------------

func Eliminar(c User) error {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("DELETE FROM users WHERE id = ?")
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
