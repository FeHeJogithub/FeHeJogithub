package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := sql.Open("mysql", dsn)

	if err != nil {

		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Conexion a la base de datos Mysql exitosamente")
	return db, nil
}
