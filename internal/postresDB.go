package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func OpenConnection() *sql.DB{

	username := "postgres"
	password := "qwerty"
	dbName := "postgres"
	dbHost := "localhost"
	dbPort := "5436"


	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password) //Создать строку подключения
	fmt.Println(dbUri)

	DB, err := sql.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	err = DB.Ping()
	if err != nil{
		panic(err)
	}
	return DB
}
