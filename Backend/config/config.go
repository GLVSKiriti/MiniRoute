package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDb() *sql.DB {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalln("Error loading .env file")
	// }

	host := os.Getenv("host")
	port := os.Getenv("port")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var db, err2 = sql.Open("postgres", psqlInfo)
	if err2 != nil {
		log.Fatalln("There's an error in connecting database")
	}

	return db
}
