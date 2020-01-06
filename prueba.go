package main

import (
	"fmt"
	"database/sql"
	_"mysql-master"
)

func main() {
	db,err:= sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/prueba")
	if err!=nil{
		fmt.Println("error")
	}

	dato, err2 := db.Query("select * from users")
	if err2!=nil{
		fmt.Println("error2")
	}

	for dato.Next(){
		var name, lastname string
		err3 := dato.Scan(&name, &lastname)
		if err3!=nil{
			fmt.Println("error3")
		}
		fmt.Println(name + " " + lastname)
	}
}
