package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type response struct {
	ID int64 `json:"id,omitempty"`
	Message string `json:"message. omitempyt"`
}

func CreateConnectionDb() *sql.DB {
	err:= godotenv.Load(".env")

	if err != nil{
		log.Fatal("Error failed to load env...")
	}

	db, err :=  sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	err= db.Ping()

	if err != nil{
		panic(err)
	}

	fmt.Println("Successfully Connected to db - postgres")
	return  db

}

