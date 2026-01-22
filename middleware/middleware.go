package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	model "go-postgrsql/models"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
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

func CreateStock(w http.ResponseWriter, r *http.Request){
	var stock model.Stock

	err:= json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatal("Unable ot decode the request body %v", err)
	}

	insertID:= insertStock(stock)
	res:= response{
		ID: insertID,
		Message: "Stock created successfully",
	}

	json.NewEncoder(w).Encode(res)
}


func GetStock(w http.ResponseWriter, r *http.Request){
	params:= mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil  {
		log.Fatal("Unable to convert the string into int. %v", err)
	}

	stock, err:= getStock(int64(id))

	if err != nil{
		log.Fatal("Unable to get stock. %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}




