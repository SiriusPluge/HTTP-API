package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

func OpenConnection() *sql.DB{
	e := godotenv.Load() //Download .env
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")


	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Создать строку подключения
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
